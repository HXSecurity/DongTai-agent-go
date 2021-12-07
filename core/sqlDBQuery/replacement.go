package sqlDBQuery

import (
	"database/sql"
	"go-agent/global"
	"go-agent/model/request"
	"go-agent/utils"
)

func Query(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	signature, callerClass, callerMethod, callerLineNumber := utils.FmtStack()
	rows, sql := QueryT(db, query, args)
	var sourceHash global.HashKeys = []string{utils.GetSource(query)}
	var SourceValues string = "query"
	var ARGS = utils.StringAdd(utils.Strval(query), ",", utils.Strval(args))
	for _, v := range args {
		switch v.(type) {
		case string:
			sourceHash = append(sourceHash, utils.GetSource(v))
			SourceValues += utils.StringAdd(v.(string), " ")
		}
	}

	var targetHash global.HashKeys = []string{}

	var pool = request.Pool{
		Source:           false,
		Interfaces:       []interface{}{},
		SourceValues:     SourceValues,
		SourceHash:       sourceHash,
		TargetValues:     "",
		TargetHash:       targetHash,
		Signature:        signature,
		OriginClassName:  "sql.(*DB)",
		MethodName:       "Query",
		ClassName:        "sql.(*DB)",
		CallerLineNumber: callerLineNumber,
		CallerClass:      callerClass,
		CallerMethod:     callerMethod,
		Args:             ARGS,
		RetClassName:     "*sql.Rows, error",
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
