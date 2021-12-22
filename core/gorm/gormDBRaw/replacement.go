package gormDBRaw

import (
	"go-agent/model/request"
	"go-agent/utils"
	"gorm.io/gorm"
)

func Raw(db *gorm.DB, sql string, values ...interface{}) (tx *gorm.DB) {
	s := RawT(db, sql, values...)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(sql, values),
		Reqs:            utils.Collect(tx),
		Source:          false,
		OriginClassName: "gorm.(*DB)",
		MethodName:      "Raw",
		ClassName:       "gorm",
	})
	return s
}

func RawT(db *gorm.DB, sql string, values ...interface{}) (tx *gorm.DB) {
	return
}
