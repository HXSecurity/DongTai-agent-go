package gormDBGroup

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"gorm.io/gorm"
)

func init() {
	model.HookMap["gormDBGroup"] = new(GormDBGroup)
}

type GormDBGroup struct {
}

func (h *GormDBGroup) Hook() {
	db := &gorm.DB{}
	err := gohook.HookMethod(db, "Group", Group, GroupT)
	if err != nil {
		fmt.Println(err, "gormGroup")
	} else {
		fmt.Println("gormGroup")
	}
}

func (h *GormDBGroup) UnHook() {
	db := &gorm.DB{}
	gohook.UnHookMethod(db, "Group")
}
