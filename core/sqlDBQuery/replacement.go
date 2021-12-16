package sqlDBQuery

import (
	"database/sql"
	"go-agent/model/request"
	"go-agent/utils"
)

func Query(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	rows, sql := QueryT(db, query, args)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(query, args),
		Reqs:            utils.Collect(rows, sql),
		Source:          false,
		OriginClassName: "sql.(*DB)",
		MethodName:      "Query",
		ClassName:       "sql.(*DB)",
	})

	return rows, sql
}

func QueryT(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}
