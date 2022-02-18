package sqlDBQuery

import (
	"database/sql"
	"github.com/HXSecurity/DongTai-agent-go/model/request"
)

func Query(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	rows, sql := QueryT(db, query, args)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(query, args),
		Reqs:            request.Collect(rows, sql),
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
