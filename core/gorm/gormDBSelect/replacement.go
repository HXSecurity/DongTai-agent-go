package gormDBSelect

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"gorm.io/gorm"
)

func Select(db *gorm.DB, query interface{}, args ...interface{}) (tx *gorm.DB) {
	s := SelectT(db, query, args...)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(query, args),
		Reqs:            request.Collect(tx),
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
