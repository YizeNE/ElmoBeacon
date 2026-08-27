package service

import (
	"ElmoBeacon/pb"
	"ElmoBeacon/util"
	"os"

	"github.com/rs/zerolog/log"
)

func GetItemDataMap(gameDataDir string, gameServer string) (map[int64]*pb.Item, error) {
	var itemData pb.ItemData
	if gameServer == string(GameServerCN) {
		err := util.GetTableData(gameDataDir, "", &itemData)
		if err != nil {
			log.Error().Err(err).Str("server", gameServer).Msg("failed to get table data (CN) when get item map")
			return nil, err
		}
	} else {
		err := util.GetTableData(gameDataDir, gameServer, &itemData)
		if err != nil {
			//当前目录不存在时，尝试到上级目录读取
			if os.IsNotExist(err) {
				err = util.GetTableData(gameDataDir, "", &itemData)
				if err != nil {
					log.Error().Err(err).Str("server", gameServer).Msg("failed to get fallback table data when get item map")
					return nil, err
				}
			} else {
				log.Error().Err(err).Str("server", gameServer).Msg("failed to get table data when get item map")
				return nil, err
			}
		}
	}

	itemMap := make(map[int64]*pb.Item)
	for i, unit := range itemData.Units {
		itemMap[unit.Id] = itemData.Units[i]
	}

	return itemMap, nil
}
