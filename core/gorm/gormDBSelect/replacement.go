package gormDBSelect

import (
	"go-agent/model/request"
	"go-agent/utils"
	"gorm.io/gorm"
)

func Select(db *gorm.DB, query interface{}, args ...interface{}) (tx *gorm.DB) {
	s := SelectT(db, query, args...)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(query, args),
		Reqs:            utils.Collect(tx),
		Source:          false,
		OriginClassName: "gorm.(*DB)",
		MethodName:      "Select",
		ClassName:       "gorm",
	})
	return s
}

func SelectT(db *gorm.DB, query interface{}, args ...interface{}) (tx *gorm.DB) {
	return
}
