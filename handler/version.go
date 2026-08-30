package handler

import (
	"ElmoBeacon/request"
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
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

// extractExeFromZip 从 zip 中提取 ElmoBeacon/ElmoBeacon.exe 并返回其内容
func extractExeFromZip(zipData []byte) ([]byte, error) {
	reader, err := zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
	if err != nil {
		return nil, err
	}

	for _, file := range reader.File {
		if strings.HasSuffix(file.Name, "ElmoBeacon.exe") && !file.FileInfo().IsDir() {
			rc, err := file.Open()
			if err != nil {
				return nil, err
			}
			defer rc.Close()
			return io.ReadAll(rc)
		}
	}

	return nil, errors.New("ElmoBeacon.exe not found in zip")
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

	downloadURL := fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/ElmoBeacon.zip", Owner, Repo, version)

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

	// 先将 zip 全部读入内存
	zipData, err := io.ReadAll(pr)
	if err != nil {
		if dlCtx.Err() == context.Canceled {
			log.Info().Msg("update canceled by user when update")
			return ErrUpdateCanceled
		}
		log.Error().Err(err).Msg("failed to read response body when update")
		return err
	}

	// 从 zip 中提取 exe
	exeData, err := extractExeFromZip(zipData)
	if err != nil {
		log.Error().Err(err).Msg("failed to extract exe from zip when update")
		return err
	}

	// 通知前端：下载完成，正在安装
	runtime.EventsEmit(a.ctx, "update:progress", map[string]interface{}{
		"percent":    100,
		"installing": true,
	})

	// 应用更新
	err = update.Apply(bytes.NewReader(exeData), update.Options{})
	if err != nil {
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
