package gormDBHaving

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
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