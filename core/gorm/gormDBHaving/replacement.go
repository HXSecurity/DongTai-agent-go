package gormDBHaving

import (
	"go-agent/model/request"
	"go-agent/utils"
	"gorm.io/gorm"
)

func Having(db *gorm.DB, query interface{}, args ...interface{}) (tx *gorm.DB) {
	s := HavingT(db, query, args...)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(query, args),
		Reqs:            utils.Collect(tx),
		Source:          false,
		OriginClassName: "gorm.(*DB)",
		MethodName:      "Having",
		ClassName:       "gorm",
	})
	return s
}

func HavingT(db *gorm.DB, query interface{}, args ...interface{}) (tx *gorm.DB) {
	return
}
