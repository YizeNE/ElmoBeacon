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

func GetLatestVersion() (string, error) {
	client, err := NewHttpClient()
	if err != nil {
		return "", err
	}

	resp, err := client.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", Owner, Repo))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New(resp.Status)
	}

	// 解析 JSON 获取 tag_name
	var release struct {
		TagName string `json:"tag_name"`
	}
	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		return "", err
	}

	return release.TagName, nil
}
