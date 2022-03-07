package gormDBExec

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"gorm.io/gorm"
)

func Exec(db *gorm.DB, sql string, values ...interface{}) (tx *gorm.DB) {
	s := ExecT(db, sql, values...)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(sql, values),
		Reqs:            request.Collect(tx),
		Source:          false,
		OriginClassName: "gorm.(*DB)",
		MethodName:      "Exec",
		ClassName:       "gorm",
	})
	return s
}

func ExecT(db *gorm.DB, sql string, values ...interface{}) (tx *gorm.DB) {
	return
}
