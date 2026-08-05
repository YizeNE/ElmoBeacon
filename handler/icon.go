package handler

import (
	"ElmoBeacon/request"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
)

// 本地缓存根目录：可执行文件所在目录/Assets
func getAssetsDir() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(filepath.Dir(exePath), "Assets")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return dir, nil
}

// 根据 iconName 判断子目录
// 角色图标: Avatar_Head_xxxUP.png → dolls/
// 武器图标: xxx_256.png           → weapons/
func getIconSubDir(iconName string) string {
	if strings.HasPrefix(iconName, "Avatar_Bust_") {
		return "dolls"
	}
	return "weapons"
}

// GetIcon 接收图标文件名，返回 base64 data URI
func (a *App) GetIcon(iconName string) (string, error) {
	if iconName == "" {
		return "", nil
	}

	subDir := getIconSubDir(iconName)

	assetsDir, err := getAssetsDir()
	if err != nil {
		return "", err
	}

	// 确保 doll/weapon 子目录存在
	iconDir := filepath.Join(assetsDir, subDir)
	if err := os.MkdirAll(iconDir, 0755); err != nil {
		return "", err
	}

	localPath := filepath.Join(iconDir, iconName)

	// 1. 本地有缓存，直接读取
	if _, err := os.Stat(localPath); err == nil {
		data, err := os.ReadFile(localPath)
		if err != nil {
			return "", err
		}
		return "data:image/png;base64," + base64.StdEncoding.EncodeToString(data), nil
	}

	// 2. 本地没有，从 GitHub 下载
	//    仓库目录结构: Assets/doll/xxx.png  Assets/weapon/xxx.png
	downloadURL := fmt.Sprintf(
		"https://cdn.jsdelivr.net/gh/YizeNE/ElmoBeaconAssets@main/%s/%s",
		subDir, iconName,
	)

	client, err := request.NewHttpClient()
	if err != nil {
		return "", err
	}

	resp, err := client.Get(downloadURL)
	if err != nil {
		log.Error().Err(err).Str("icon", iconName).Msg("failed to download icon")
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Warn().Str("icon", iconName).Msg("icon not found")
		return "", nil
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 3. 存到本地缓存
	if err := os.WriteFile(localPath, data, 0644); err != nil {
		log.Error().Err(err).Msg("failed to cache icon locally")
	}

	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(data), nil
}
