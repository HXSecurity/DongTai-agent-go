package gormDBGroup

import (
	"go-agent/model/request"
	"go-agent/utils"
	"gorm.io/gorm"
)

func Group(db *gorm.DB, value string) (tx *gorm.DB) {
	s := GroupT(db, value)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(value),
		Reqs:            utils.Collect(tx),
		Source:          false,
		OriginClassName: "gorm.(*DB)",
		MethodName:      "Group",
		ClassName:       "gorm",
	})
	return s
}

func GroupT(db *gorm.DB, value string) (tx *gorm.DB) {
	return
}
