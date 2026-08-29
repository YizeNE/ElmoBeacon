package service

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

const (
	gameLogPathCN      = "AppData/LocalLow/SunBorn/少女前线2：追放/Player.log"
	gameLogPathOversea = "AppData/LocalLow/SunBorn/EXILIUM/Player.log"
)

const (
	exprGameDataDir    = `\[Subsystems] Discovering subsystems at path (.+)/UnitySubsystems`
	exprGachaRecordUrl = `POST\s*(https://[^\s]+/list[^\s]*)`
	exprLoginInfo      = `Authorization:\s*([^\s]+)`
)

type gameServer string

const (
	GameServerCN     gameServer = "cn"   //DarkWinter China
	GameServerUS     gameServer = "us"   //DarkWinter USA
	GameServerGlobal gameServer = "intl" //HaoPlay Global
	GameServerJP     gameServer = "jp"   //HaoPlay Japan
	GameServerKR     gameServer = "kr"   //HaoPlay Korea
	GameServerAsia   gameServer = "tw"   //HaoPlay Asia
)

type GameUserInfo struct {
	Uid             uint64
	GameServer      gameServer
	GameDataDir     string
	GameAccessToken string
	GachaRecordUrl  string
}

func GetGameDataDir(server gameServer) (string, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		log.Error().Err(err).Msg("failed to get user home dir when get game data dir")
		return "", errors.New("Failed to get user home dir")
	}

	var logPath string
	if server == GameServerCN {
		logPath = filepath.Join(userHome, gameLogPathCN)
	} else {
		logPath = filepath.Join(userHome, gameLogPathOversea)
	}

	logDataBytes, err := os.ReadFile(logPath)
	if err != nil {
		log.Error().Err(err).Str("server", string(server)).Msg("failed to read game log file when get game data dir")
		return "", errors.Errorf("Failed to read game log file(%s)", server)
	}

	regexpGameDataDir, err := regexp.Compile(exprGameDataDir)
	if err != nil {
		log.Error().Err(err).Msg("failed to compile exprGameDataDir when get game data dir")
		return "", errors.New("Failed to compile exprGameDataDir")
	}

	resultGameDataDir := regexpGameDataDir.FindSubmatch(logDataBytes)
	if resultGameDataDir == nil {
		log.Error().Str("server", string(server)).Msg("failed to find game data directory when get game data dir")
		return "", errors.Errorf("Failed to find game data directory(%s)", server)
	}

	return filepath.Join(string(resultGameDataDir[1]), "LocalCache/Data"), nil
}

func GetUserInfo(uid uint64, filePath string) (*GameUserInfo, error) {
	captureDataBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Error().Err(err).Str("filePath", filePath).Msg("failed to read capture file when get user info")
		return nil, errors.New("Failed to read capture file")
	}

	//extract GachaRecordUrl
	regexpGachaRecordUrl, err := regexp.Compile(exprGachaRecordUrl)
	if err != nil {
		log.Error().Err(err).Msg("failed to compile exprGachaRecordUrl when get user info")
		return nil, errors.New("Failed to compile exprGachaRecordUrl")
	}
	resultGachaRecordUrlList := regexpGachaRecordUrl.FindSubmatch(captureDataBytes)
	if resultGachaRecordUrlList == nil {
		log.Error().Msg("failed to find gacha record url when get user info")
		return nil, errors.New("Failed to find gacha record url")
	}
	gachaRecordUrl := string(resultGachaRecordUrlList[len(resultGachaRecordUrlList)-1])

	//determine the server
	var server gameServer
	switch {
	case strings.Contains(gachaRecordUrl, "gf2-gacha-record-us"):
		server = GameServerUS
	case strings.Contains(gachaRecordUrl, "gf2-gacha-record-intl"):
		server = GameServerGlobal
	case strings.Contains(gachaRecordUrl, "gf2-gacha-record-jp"):
		server = GameServerJP
	case strings.Contains(gachaRecordUrl, "gf2-gacha-record-kr"):
		server = GameServerKR
	case strings.Contains(gachaRecordUrl, "gf2-gacha-record-asia"):
		server = GameServerAsia
	case strings.Contains(gachaRecordUrl, "gf2-gacha-record"):
		server = GameServerCN
	default:
		log.Error().Str("url", gachaRecordUrl).Msg("failed to determine server when get user info")
		return nil, errors.Errorf("Failed to determine server,gacha url:%s", gachaRecordUrl)
	}

	//extract accessToken
	regexpLoginInfo, err := regexp.Compile(exprLoginInfo)
	if err != nil {
		log.Error().Err(err).Msg("failed to compile exprLoginInfo when get user info")
		return nil, errors.New("Failed to compile exprLoginInfo")
	}
	resultLoginInfoList := regexpLoginInfo.FindSubmatch(captureDataBytes)
	if resultLoginInfoList == nil {
		log.Error().Msg("failed to find game login information when get user info")
		return nil, errors.New("Failed to find game login information")
	}
	gameAccessToken := string(resultLoginInfoList[len(resultLoginInfoList)-1])

	gameDataDir, err := GetGameDataDir(server)
	if err != nil {
		return nil, err
	}

	return &GameUserInfo{
		Uid:             uid,
		GameServer:      server,
		GameDataDir:     gameDataDir,
		GameAccessToken: gameAccessToken,
		GachaRecordUrl:  gachaRecordUrl,
	}, nil
}
