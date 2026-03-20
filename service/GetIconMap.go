package service

import (
	"ElmoBeacon/pb"
	"ElmoBeacon/util"
	"os"

	"github.com/gookit/color"
)

func GetDollIconMap(gameDataDir string, gameServer string) (map[int64]string, error) {
	var gunData pb.GunData
	if gameServer == string(GameServerCN) {
		err := util.GetTableData(gameDataDir, "", &gunData)
		if err != nil {
			return nil, err
		}
	} else {
		err := util.GetTableData(gameDataDir, gameServer, &gunData)
		if err != nil {
			//当前目录不存在时，尝试到上级目录读取
			if os.IsNotExist(err) {
				err = util.GetTableData(gameDataDir, "", &gunData)
				if err != nil {
					return nil, err
				}
			} else {
				return nil, err
			}
		}
	}

	iconMap := make(map[int64]string)
	for i, unit := range gunData.Units {
		iconMap[unit.Id] = color.Sprintf("Avatar_Head_%sUP.png", gunData.Units[i].Code)
	}

	return iconMap, nil
}

func GetWeaponIconMap(gameDataDir string, gameServer string) (map[int64]string, error) {
	var weaponData pb.GunWeaponData
	if gameServer == string(GameServerCN) {
		err := util.GetTableData(gameDataDir, "", &weaponData)
		if err != nil {
			return nil, err
		}
	} else {
		err := util.GetTableData(gameDataDir, gameServer, &weaponData)
		if err != nil {
			//当前目录不存在时，尝试到上级目录读取
			if os.IsNotExist(err) {
				err = util.GetTableData(gameDataDir, "", &weaponData)
				if err != nil {
					return nil, err
				}
			} else {
				return nil, err
			}
		}
	}

	iconMap := make(map[int64]string)
	for i, unit := range weaponData.Units {
		iconMap[unit.Id] = color.Sprintf("%s_256.png", weaponData.Units[i].Code)
	}

	return iconMap, nil
}
