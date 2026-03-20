package service

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pkg/errors"
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

func GetUserInfo(uid uint64, filePath string) (*GameUserInfo, error) {
	captureDataBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("Failed to read capture file")
	}

	//extract GachaRecordUrl
	regexpGachaRecordUrl, err := regexp.Compile(exprGachaRecordUrl)
	if err != nil {
		return nil, errors.New("Failed to compile exprGachaRecordUrl")
	}
	resultGachaRecordUrlList := regexpGachaRecordUrl.FindSubmatch(captureDataBytes)
	if resultGachaRecordUrlList == nil {
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
		return nil, errors.Errorf("Failed to determine server,gacha url:%s", gachaRecordUrl)
	}

	//extract accessToken
	regexpLoginInfo, err := regexp.Compile(exprLoginInfo)
	if err != nil {
		return nil, errors.New("Failed to compile exprLoginInfo")
	}
	resultLoginInfoList := regexpLoginInfo.FindSubmatch(captureDataBytes)
	if resultLoginInfoList == nil {
		return nil, errors.New("Failed to find game login information")
	}
	gameAccessToken := string(resultLoginInfoList[len(resultLoginInfoList)-1])

	userHome, err := os.UserHomeDir()
	if err != nil {
		return nil, errors.New("Failed to get user home dir")
	}

	//extract GameDataDir
	var gameDataDir string
	if server == GameServerCN {
		logDataBytes, err := os.ReadFile(filepath.Join(userHome, gameLogPathCN))
		if err != nil {
			return nil, errors.New("Failed to read game log file(CN)")
		}

		regexpGameDataDir, err := regexp.Compile(exprGameDataDir)
		if err != nil {
			return nil, errors.New("Failed to compile exprGameDataDir(CN)")
		}
		resultGameDataDir := regexpGameDataDir.FindSubmatch(logDataBytes)
		if resultGameDataDir == nil {
			return nil, errors.New("Failed to find game data directory(CN)")
		}
		gameDataDir = filepath.Join(string(resultGameDataDir[1]), "LocalCache/Data")
	} else {
		logDataBytes, err := os.ReadFile(filepath.Join(userHome, gameLogPathOversea))
		if err != nil {
			return nil, errors.New("Failed to read game log file(Oversea)")
		}

		regexpGameDataDir, err := regexp.Compile(exprGameDataDir)
		if err != nil {
			return nil, errors.New("Failed to compile exprGameDataDir(Oversea)")
		}
		resultGameDataDir := regexpGameDataDir.FindSubmatch(logDataBytes)
		if resultGameDataDir == nil {
			return nil, errors.New("Failed to find game data directory(Oversea)")
		}
		gameDataDir = filepath.Join(string(resultGameDataDir[1]), "LocalCache/Data")
	}

	return &GameUserInfo{
		Uid:             uid,
		GameServer:      server,
		GameDataDir:     gameDataDir,
		GameAccessToken: gameAccessToken,
		GachaRecordUrl:  gachaRecordUrl,
	}, nil
}
