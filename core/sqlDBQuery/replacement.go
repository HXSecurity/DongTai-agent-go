package sqlDBQuery

import (
	"database/sql"
	"go-agent/global"
	"go-agent/model/request"
	"go-agent/utils"
)

func Query(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {

	rows, sql := QueryT(db, query, args)

	var sourceHash global.HashKeys = []string{utils.GetSource(query)}
	var sourceBytes = []byte("")
	for _, v := range args {
		switch v.(type) {
		case string:
			sourceHash = append(sourceHash, utils.GetSource(v))
			space := []byte(" ")
			str1 := append([]byte(v.(string)), space...)
			sourceBytes = append(sourceBytes, str1...)
		}
	}

	var targetHash global.HashKeys = []string{}
	var pool = request.Pool{
		SourceValues: string(sourceBytes),
		SourceHash:   sourceHash,
		TargetHash:   targetHash,
	}

	poolTree := request.PoolTree{
		Begin:       false,
		Pool:        &pool,
		Children:    []*request.PoolTree{},
		GoroutineID: utils.CatGoroutineID(),
	}
	for k, _ := range global.PoolTreeMap {
		if k.Some(sourceHash) {
			global.PoolTreeMap[k].Children = append(global.PoolTreeMap[k].Children, &poolTree)
			break
		}
	}

	global.PoolTreeMap[&targetHash] = &poolTree

	return rows, sql
}

func QueryT(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}
