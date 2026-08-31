package service

import (
	"ElmoBeacon/db"
	"ElmoBeacon/model"
	"fmt"
	"slices"
	"strconv"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var ServerKeyMap = map[string]string{
	"hao-jp":   "jp",
	"hao-kr":   "kr",
	"hao-asia": "tw",
	"hao-intl": "intl",
	"dw-cn":    "cn",
	"dw-us":    "us",
}

type ImportDiff struct {
	PoolType     int64
	SuccessCount int
	FailCount    int
}

type ImportResult struct {
	Id       int64
	Server   string
	Uid      uint64
	DiffList []ImportDiff
}

type ImportRecord struct {
	PoolId  int64 `json:"pool_id"`
	ItemId  int64 `json:"item"`
	Time    int64 `json:"time"`
	ItemNum int64 `json:"item_num"`
	Quality int64 `json:"quality"`
}

type ImportFile map[string]map[string][]ImportRecord

func ImportRecords(uid uint64, server string, importData ImportFile) (*ImportResult, error) {
	importResult := ImportResult{
		Server: server,
		Uid:    uid,
	}

	cond := model.User{GameServer: server, Uid: uid}
	hasUser, err := db.Engine.Get(&cond)
	if err != nil {
		log.Error().Err(err).Msg("failed to query user when import records")
		return nil, errors.New("Failed to query user")
	}

	var userId int64
	if !hasUser {
		gameDataDir, err := GetGameDataDir(gameServer(server))
		if err != nil {
			return nil, err
		}
		user := model.User{Uid: uid, GameServer: server, GameDataDir: gameDataDir}
		_, err = db.Engine.Insert(&user)
		if err != nil {
			log.Error().Err(err).Msg("failed to insert user when import records")
			return nil, errors.New("Failed to insert user")
		}
		userId = user.Id
	} else {
		userId = cond.Id
	}
	importResult.Id = userId

	var serverKey string
	for k, v := range ServerKeyMap {
		if v == server {
			serverKey = k
			break
		}
	}
	if serverKey == "" {
		log.Error().Str("server", server).Msg("unsupported server when import records")
		return nil, errors.New("Unsupported server")
	}

	poolMap, ok := importData[serverKey]
	if !ok {
		log.Error().Str("serverKey", serverKey).Msg("no records found for this server in import file")
		return nil, errors.New("No records found for this server in import file")
	}

	var poolTypeKeys []string
	for k := range poolMap {
		poolTypeKeys = append(poolTypeKeys, k)
	}
	slices.Sort(poolTypeKeys)
	for _, poolTypeStr := range poolTypeKeys {
		importRecords := poolMap[poolTypeStr]
		poolType, _ := strconv.ParseInt(poolTypeStr, 10, 64)

		var existingRecords []model.Record
		err = db.Engine.Where("user_id = ? AND pool_type = ?", userId, poolType).
			Asc("id").Find(&existingRecords)
		if err != nil {
			log.Error().Err(err).Int64("poolType", poolType).Msg("failed to query existing records when import records")
			return nil, errors.New("Failed to query existing records")
		}

		var importModels []model.Record
		for _, ir := range importRecords {
			importModels = append(importModels, model.Record{
				UserId:    userId,
				PoolType:  poolType,
				PoolId:    ir.PoolId,
				ItemId:    ir.ItemId,
				Timestamp: ir.Time,
				Rank:      ir.Quality,
			})
		}
		slices.Reverse(importModels)

		merged, successCount, failCount := MergeByTimestamp(existingRecords, importModels, userId, poolType)

		ReassignVirtualIds(merged)

		session := db.Engine.NewSession()
		if err := session.Begin(); err != nil {
			log.Error().Err(err).Msg("failed to begin transaction when import records")
			return nil, errors.New("Failed to begin transaction")
		}
		if _, err := session.Where("user_id = ? AND pool_type = ?", userId, poolType).Delete(&model.Record{}); err != nil {
			session.Rollback()
			log.Error().Err(err).Msg("failed to delete old records when import records")
			return nil, errors.New("Failed to delete old records")
		}
		if len(merged) > 0 {
			batchSize := 100
			for i := 0; i < len(merged); i += batchSize {
				end := i + batchSize
				if end > len(merged) {
					end = len(merged)
				}
				batch := merged[i:end]
				if _, err := session.Insert(&batch); err != nil {
					session.Rollback()
					log.Error().Err(err).Msg("failed to insert merged records when import records")
					return nil, errors.New("Failed to insert merged records")
				}
			}
		}
		if err := session.Commit(); err != nil {
			log.Error().Err(err).Msg("failed to commit transaction when import records")
			return nil, errors.New("Failed to commit transaction")
		}
		session.Close()
		importResult.DiffList = append(importResult.DiffList, ImportDiff{
			PoolType:     poolType,
			SuccessCount: successCount,
			FailCount:    failCount,
		})

		log.Info().Int64("poolType", poolType).Int("success", successCount).Int("fail", failCount).Int("merged", len(merged)).Msg("import merged")
	}

	return &importResult, nil
}

func MergeByTimestamp(existing, imported []model.Record, userId int64, poolType int64) ([]model.Record, int, int) {
	existingByTs := groupByTimestamp(existing)
	importedByTs := groupByTimestamp(imported)
	allTimestamps := sortedUnionKeys(existingByTs, importedByTs)

	var merged []model.Record
	var successCount, failCount int

	for _, ts := range allTimestamps {
		E := existingByTs[ts]
		I := importedByTs[ts]

		switch {
		case len(E) == 0:
			successCount += len(I)
			merged = append(merged, I...)
		case len(I) == 0:
			merged = append(merged, E...)
		default:
			common, newRec, fc, conflict := AlignSequence(E, I)
			if conflict != "" {
				log.Warn().Int64("poolType", poolType).Int64("timestamp", ts).Int("failCount", fc).Msg(conflict)
				failCount += fc
				merged = append(merged, E...)
			} else {
				successCount += len(newRec)
				merged = append(merged, common...)
				merged = append(merged, newRec...)
			}
		}
	}
	return merged, successCount, failCount
}

func AlignSequence(existing, imported []model.Record) ([]model.Record, []model.Record, int, string) {
	minLen := len(existing)
	if len(imported) < minLen {
		minLen = len(imported)
	}

	divergeAt := -1
	for i := 0; i < minLen; i++ {
		if existing[i].ItemId != imported[i].ItemId {
			divergeAt = i
			break
		}
	}

	if divergeAt == -1 {
		if len(existing) <= len(imported) {
			return imported[:len(existing)], imported[len(existing):], 0, ""
		}
		return existing, nil, 0, ""
	}

	return nil, nil, len(imported),
		fmt.Sprintf("diverge at position %d: existing item=%d, imported item=%d", divergeAt, existing[divergeAt].ItemId, imported[divergeAt].ItemId)
}

func ReassignVirtualIds(records []model.Record) {
	var lastTs int64
	var order int64
	for i := range records {
		if records[i].Timestamp != lastTs {
			order = 0
			lastTs = records[i].Timestamp
		}
		records[i].Id, _ = strconv.ParseUint(fmt.Sprintf("%d%03d", records[i].Timestamp, order), 10, 64)
		order++
	}
}

func groupByTimestamp(records []model.Record) map[int64][]model.Record {
	m := make(map[int64][]model.Record)
	for _, r := range records {
		m[r.Timestamp] = append(m[r.Timestamp], r)
	}
	return m
}

func sortedUnionKeys(a, b map[int64][]model.Record) []int64 {
	seen := make(map[int64]bool)
	var keys []int64
	for k := range a {
		keys = append(keys, k)
		seen[k] = true
	}
	for k := range b {
		if !seen[k] {
			keys = append(keys, k)
		}
	}
	slices.Sort(keys)
	return keys
}
