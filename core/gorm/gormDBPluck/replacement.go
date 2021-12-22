package gormDBPluck

import (
	"go-agent/model/request"
	"go-agent/utils"
	"gorm.io/gorm"
)

func Pluck(db *gorm.DB, column string, dest interface{}) (tx *gorm.DB) {
	s := PluckT(db, column, dest)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(column, dest),
		Reqs:            utils.Collect(tx),
		Source:          false,
		OriginClassName: "gorm.(*DB)",
		MethodName:      "Pluck",
		ClassName:       "gorm",
	})
	return s
}

func PluckT(db *gorm.DB, column string, dest interface{}) (tx *gorm.DB) {
	return
}
