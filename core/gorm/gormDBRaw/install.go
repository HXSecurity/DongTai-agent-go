package gormDBRaw

import (
	"fmt"
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
	"gorm.io/gorm"
)

func init() {
	model.HookMap["gormDBRaw"] = new(GormDBRaw)
}

type GormDBRaw struct {
}

func (h *GormDBRaw) Hook() {
	db := &gorm.DB{}
	err := gohook.HookMethod(db, "Raw", Raw, RawT)
	if err != nil {
		fmt.Println(err, "gormRaw")
	} else {
		fmt.Println("gormRaw")
	}
}

func (h *GormDBRaw) UnHook() {
	db := &gorm.DB{}
	gohook.UnHookMethod(db, "Raw")
}
