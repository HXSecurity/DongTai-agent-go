package gormDBExec

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"gorm.io/gorm"
)

func Exec(db *gorm.DB, sql string, values ...interface{}) (tx *gorm.DB) {
	s := ExecT(db, sql, values...)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(sql, values),
		Reqs:            utils.Collect(tx),
		Source:          false,
		OriginClassName: "gorm.(*DB)",
		MethodName:      "Exec",
		ClassName:       "gorm",
	})
	return s
}

func ExecT(db *gorm.DB, sql string, values ...interface{}) (tx *gorm.DB) {
	return
}
