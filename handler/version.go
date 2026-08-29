package handler

import (
	"ElmoBeacon/request"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"sync"

	"github.com/inconshreveable/go-update"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const (
	Owner = "YizeNE"
	Repo  = "ElmoBeacon"
)

const Version = ""

// updateMu 保护 updateCancel 的并发安全
var (
	updateMu     sync.Mutex
	updateCancel context.CancelFunc
)

// ErrUpdateCanceled 用户主动取消更新
var ErrUpdateCanceled = errors.New("update canceled by user")

// progressReader 包装 io.Reader，每读取一段数据就向前端推送进度
type progressReader struct {
	reader  io.Reader
	total   int64
	read    int64
	ctx     context.Context
	lastPct int // 上次推送的百分比，避免过于频繁推送
}

func (pr *progressReader) Read(p []byte) (int, error) {
	n, err := pr.reader.Read(p)
	if n > 0 {
		pr.read += int64(n)
		if pr.total > 0 {
			pct := int(float64(pr.read) / float64(pr.total) * 100)
			// 每变化 1% 推送一次，避免事件洪泛
			if pct != pr.lastPct {
				pr.lastPct = pct
				runtime.EventsEmit(pr.ctx, "update:progress", map[string]interface{}{
					"downloaded": pr.read,
					"total":      pr.total,
					"percent":    pct,
				})
			}
		}
	}
	return n, err
}

func (a *App) GetVersion() string {
	return Version
}

func (a *App) GetLatestRelease() (*request.ReleaseInfo, error) {
	release, err := request.GetLatestRelease()
	if err != nil {
		log.Error().Err(err).Msg("failed to get latest release when check update")
		return nil, err
	}
	return release, nil
}

func (a *App) UpdateTo(version string) error {
	client, err := request.NewDownloadClient()
	if err != nil {
		log.Error().Err(err).Msg("failed to create download client when update")
		return err
	}

	downloadURL := fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/ElmoBeacon.exe", Owner, Repo, version)

	// 创建可取消的 context
	dlCtx, cancel := context.WithCancel(context.Background())
	updateMu.Lock()
	updateCancel = cancel
	updateMu.Unlock()

	// 确保 UpdateTo 退出时清理 updateCancel
	defer func() {
		updateMu.Lock()
		updateCancel = nil
		updateMu.Unlock()
	}()

	req, err := http.NewRequestWithContext(dlCtx, http.MethodGet, downloadURL, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to create request when update")
		cancel()
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		// 用户取消，返回特定错误
		if dlCtx.Err() == context.Canceled {
			log.Info().Msg("update canceled by user when update")
			return ErrUpdateCanceled
		}
		log.Error().Err(err).Str("url", downloadURL).Msg("failed to download when update")
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error().Int("status", resp.StatusCode).Str("url", downloadURL).Msg("unexpected status code when update")
		return errors.New(resp.Status)
	}

	// 用 progressReader 包装响应体，边下载边推送进度
	pr := &progressReader{
		reader: resp.Body,
		total:  resp.ContentLength,
		ctx:    a.ctx, // Wails app context，用于 EventsEmit
	}

	err = update.Apply(pr, update.Options{})
	if err != nil {
		if dlCtx.Err() == context.Canceled {
			log.Info().Msg("update canceled by user when update")
			return ErrUpdateCanceled
		}
		log.Error().Err(err).Msg("failed to apply update when update")
		return err
	}

	execPath, err := os.Executable()
	if err != nil {
		log.Error().Err(err).Msg("failed to get executable path when update")
		return err
	}
	cmd := exec.Command(execPath)
	err = cmd.Start()
	if err != nil {
		log.Error().Err(err).Msg("failed to restart process when update")
		return err
	}
	os.Exit(0)

	return nil
}

// CancelUpdate 供前端调用，取消正在进行的下载
func (a *App) CancelUpdate() {
	updateMu.Lock()
	defer updateMu.Unlock()
	if updateCancel != nil {
		updateCancel()
		updateCancel = nil
	}
}
