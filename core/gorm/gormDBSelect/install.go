package gormDBSelect

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"gorm.io/gorm"
)

func init() {
	model.HookMap["gormDBSelect"] = new(GormDBSelect)
}

type GormDBSelect struct {
}

func (h *GormDBSelect) Hook() {
	db := &gorm.DB{}
	err := gohook.HookMethod(db, "Select", Select, SelectT)
	if err != nil {
		fmt.Println(err, "gormOrder")
	} else {
		fmt.Println("gormOrder")
	}
}

func (h *GormDBSelect) UnHook() {
	db := &gorm.DB{}
	gohook.UnHookMethod(db, "Select")
}
