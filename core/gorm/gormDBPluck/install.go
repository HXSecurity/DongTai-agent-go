package gormDBPluck

import (
	"fmt"
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
	"gorm.io/gorm"
)

func init() {
	model.HookMap["gormDBPluck"] = new(GormDBPluck)
}

type GormDBPluck struct {
}

func (h *GormDBPluck) Hook() {
	db := &gorm.DB{}
	err := gohook.HookMethod(db, "Pluck", Pluck, PluckT)
	if err != nil {
		fmt.Println(err, "GormDBPluck")
	} else {
		fmt.Println("GormDBPluck")
	}
}

func (h *GormDBPluck) UnHook() {
	db := &gorm.DB{}
	gohook.UnHookMethod(db, "Pluck")
}
