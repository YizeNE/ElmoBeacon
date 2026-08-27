package service

import (
	"ElmoBeacon/pb"
	"ElmoBeacon/util"

	"github.com/rs/zerolog/log"
)

func GetLangDataMap(gameDataDir, gameServer, lang string) (map[int64]string, error) {
	langMap := make(map[int64]string)
	switch {
	case gameServer == string(GameServerCN):
		var langData pb.LangPackageTableCnData
		err := util.GetTableData(gameDataDir, "", &langData)
		if err != nil {
			log.Error().Err(err).Str("server", gameServer).Str("lang", lang).Msg("failed to get table data when get lang map")
			return nil, err
		}
		for _, unit := range langData.Units {
			langMap[unit.Id] = unit.Content
		}
	case lang == "zh-CN":
		var langData pb.LangPackageTableCnData
		err := util.GetTableData(gameDataDir, "", &langData)
		if err != nil {
			log.Error().Err(err).Str("server", gameServer).Str("lang", lang).Msg("failed to get table data when get lang map")
			return nil, err
		}
		for _, unit := range langData.Units {
			langMap[unit.Id] = unit.Content
		}
	case lang == "zh-TW":
		var langData pb.LangPackageTableZhtcData
		err := util.GetTableData(gameDataDir, "", &langData)
		if err != nil {
			log.Error().Err(err).Str("server", gameServer).Str("lang", lang).Msg("failed to get table data when get lang map")
			return nil, err
		}
		for _, unit := range langData.Units {
			langMap[unit.Id] = unit.Content
		}
	case lang == "ja":
		var langData pb.LangPackageTableJajpData
		err := util.GetTableData(gameDataDir, "", &langData)
		if err != nil {
			log.Error().Err(err).Str("server", gameServer).Str("lang", lang).Msg("failed to get table data when get lang map")
			return nil, err
		}
		for _, unit := range langData.Units {
			langMap[unit.Id] = unit.Content
		}
	case lang == "kr":
		var langData pb.LangPackageTableKokrData
		err := util.GetTableData(gameDataDir, "", &langData)
		if err != nil {
			log.Error().Err(err).Str("server", gameServer).Str("lang", lang).Msg("failed to get table data when get lang map")
			return nil, err
		}
		for _, unit := range langData.Units {
			langMap[unit.Id] = unit.Content
		}
	default:
		var langData pb.LangPackageTableEnusData
		err := util.GetTableData(gameDataDir, "", &langData)
		if err != nil {
			log.Error().Err(err).Str("server", gameServer).Str("lang", lang).Msg("failed to get table data when get lang map")
			return nil, err
		}
		for _, unit := range langData.Units {
			langMap[unit.Id] = unit.Content
		}
	}

	return langMap, nil
}
