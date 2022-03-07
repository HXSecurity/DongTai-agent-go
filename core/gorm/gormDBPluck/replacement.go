package gormDBPluck

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"gorm.io/gorm"
)

func Pluck(db *gorm.DB, column string, dest interface{}) (tx *gorm.DB) {
	s := PluckT(db, column, dest)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(column, dest),
		Reqs:            request.Collect(tx),
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
