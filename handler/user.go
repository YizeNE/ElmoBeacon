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
		log.Error().Err(err).Msg("")
		return nil, errors.New("error occurred when query user list from db")
	}
	return userList, nil
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
		return "", err
	}
	if filePath == "" {
		return "", errors.New("未选择文件")
	}
	return filePath, nil
}
