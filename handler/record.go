package handler

import (
	"ElmoBeacon/service"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// SyncRecords 使用uid和文件路径进行同步
func (a *App) SyncRecords(uid uint64, filePath string) (*service.SyncResult, error) {
	gameUserInfo, err := service.GetUserInfo(uid, filePath)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, errors.Errorf("Failed to extract game user information")
	}

	// 调用service.SyncRecords进行同步
	syncResult, err := service.SyncRecords(gameUserInfo)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}

	return syncResult, nil
}
