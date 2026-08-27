package handler

import (
	"ElmoBeacon/request"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/inconshreveable/go-update"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

const (
	Owner = "YizeNE"
	Repo  = "ElmoBeacon"
)

const Version = ""

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
	client, err := request.NewHttpClient()
	if err != nil {
		log.Error().Err(err).Msg("failed to create http client when update")
		return err
	}
	downloadURL := fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/ElmoBeacon.exe", Owner, Repo, version)
	resp, err := client.Get(downloadURL)
	if err != nil {
		log.Error().Err(err).Str("url", downloadURL).Msg("failed to download when update")
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Error().Int("status", resp.StatusCode).Str("url", downloadURL).Msg("unexpected status code when update")
		return errors.New(resp.Status)
	}

	err = update.Apply(resp.Body, update.Options{})
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
