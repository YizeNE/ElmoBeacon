package service

import (
	"ElmoBeacon/pb"
	"ElmoBeacon/util"

	"github.com/rs/zerolog/log"
)

func GetGachaPoolTypeList(gameDataDir string) (gachaPoolTypeList []int64, err error) {
	var gachaTypeListData pb.GachaTypeListData

	err = util.GetTableData(gameDataDir, "", &gachaTypeListData)
	if err != nil {
		log.Error().Err(err).Msg("failed to get table data when get gacha pool type list")
		return nil, err
	}

	for _, unit := range gachaTypeListData.Units {
		gachaPoolTypeList = append(gachaPoolTypeList, unit.Id)
	}

	return gachaPoolTypeList, nil
}
