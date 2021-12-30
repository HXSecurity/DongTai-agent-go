package gormDBOrder

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"gorm.io/gorm"
)

func Order(db *gorm.DB, value interface{}) (tx *gorm.DB) {

	s := OrderT(db, value)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(value),
		Reqs:            utils.Collect(tx),
		Source:          false,
		OriginClassName: "gorm.(*DB)",
		MethodName:      "Order",
		ClassName:       "gorm",
	})
	return s
}

func OrderT(db *gorm.DB, value interface{}) (tx *gorm.DB) {
	return
}