package gormDBRaw

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"gorm.io/gorm"
)

func Raw(db *gorm.DB, sql string, values ...interface{}) (tx *gorm.DB) {
	s := RawT(db, sql, values...)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(sql, values),
		Reqs:            request.Collect(tx),
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
