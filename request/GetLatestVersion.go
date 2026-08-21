package request

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

const (
	Owner = "YizeNE"
	Repo  = "ElmoBeacon"
)

type ReleaseInfo struct {
	TagName string `json:"tag_name"`
	Body    string `json:"body"` // 更新日志（Markdown）
}

func GetLatestRelease() (*ReleaseInfo, error) {
	client, err := NewHttpClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", Owner, Repo))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	var release ReleaseInfo
	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		return nil, err
	}

	return &release, nil
}
