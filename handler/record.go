package handler

import (
	"ElmoBeacon/service"

	"github.com/pkg/errors"
)

// SyncRecords 使用uid和文件路径进行同步
func (a *App) SyncRecords(uid uint64, filePath string) (*service.SyncResult, error) {
	gameUserInfo, err := service.GetUserInfo(uid, filePath)
	if err != nil {
		return nil, errors.Errorf("Failed to extract game user information")
	}

	// 调用service.SyncRecords进行同步
	syncResult, err := service.SyncRecords(a.ctx, gameUserInfo)
	if err != nil {
		return nil, err
	}

	return syncResult, nil
}

// ImportRecords 导入外部抽卡记录
func (a *App) ImportRecords(uid uint64, server string, importData service.ImportFile) (*service.ImportResult, error) {
	importResult, err := service.ImportRecords(uid, server, importData)
	if err != nil {
		return nil, err
	}
	return importResult, nil
}
