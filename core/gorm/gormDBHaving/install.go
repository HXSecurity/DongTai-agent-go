package gormDBHaving

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"gorm.io/gorm"
)

func init() {
	model.HookMap["gormDBHaving"] = new(GormDBHaving)
}

type GormDBHaving struct {
}

func (h *GormDBHaving) Hook() {
	db := &gorm.DB{}
	err := gohook.HookMethod(db, "Having", Having, HavingT)
	if err != nil {
		fmt.Println(err, "gormHaving")
	} else {
		fmt.Println("gormHavingr")
	}
}

func (h *GormDBHaving) UnHook() {
	db := &gorm.DB{}
	gohook.UnHookMethod(db, "Having")
}
