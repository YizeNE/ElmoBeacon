package handler

import (
	"ElmoBeacon/db"
	"ElmoBeacon/model"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) GetUserList() ([]model.User, error) {
	var userList []model.User
	err := db.Engine.Find(&userList)
	if err != nil {
		log.Error().Err(err).Msg("failed to query user list from db when select file path")
		return nil, errors.New("error occurred when query user list from db")
	}
	return userList, nil
}

// DeleteUser 删除用户及其关联的抽卡记录
func (a *App) DeleteUser(id int64) error {
	// 先删除该用户关联的所有抽卡记录
	_, err := db.Engine.Where("user_id = ?", id).Delete(&model.Record{})
	if err != nil {
		log.Error().Err(err).Msg("")
		return errors.New("error occurred when deleting user records from db")
	}

	// 再删除用户本身
	_, err = db.Engine.ID(id).Delete(&model.User{})
	if err != nil {
		log.Error().Err(err).Msg("")
		return errors.New("error occurred when deleting user from db")
	}
	return nil
}

// SelectFilePath 只选择文件，不解析
func (a *App) SelectFilePath() (string, error) {
	filePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "选择抓包文件",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "文本文件",
				Pattern:     "*.txt",
			},
			{
				DisplayName: "所有文件",
				Pattern:     "*.*",
			},
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to open file dialog when select file path")
		return "", err
	}
	return filePath, nil
}
