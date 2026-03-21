package handler

import (
	"ElmoBeacon/request"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/inconshreveable/go-update"
	"github.com/pkg/errors"
)

const (
	Owner = "YizeNE"
	Repo  = "ElmoBeacon"
)

const Version = ""

func (a *App) GetVersion() string {
	return Version
}

func (a *App) GetLatestVersion() (string, error) {
	return request.GetLatestVersion()
}

func (a *App) UpdateTo(version string) error {
	downloadURL := fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/ElmoBeacon.exe", Owner, Repo, version)
	resp, err := http.Get(downloadURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		return err
	}

	execPath, err := os.Executable()
	if err != nil {
		return err
	}
	cmd := exec.Command(execPath)
	err = cmd.Start()
	if err != nil {
		return err
	}
	os.Exit(0)

	return nil
}
