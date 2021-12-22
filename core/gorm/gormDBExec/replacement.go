package gormDBExec

import (
	"go-agent/model/request"
	"go-agent/utils"
	"gorm.io/gorm"
)

func Exec(db *gorm.DB, sql string, values ...interface{}) (tx *gorm.DB) {
	s := ExecT(db, sql, values...)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(sql, values),
		Reqs:            utils.Collect(tx),
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
