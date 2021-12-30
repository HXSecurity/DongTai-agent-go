package gormDBPluck

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
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
