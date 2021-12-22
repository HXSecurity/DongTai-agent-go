package gormDBExec

import (
	"fmt"
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
	"gorm.io/gorm"
)

func init() {
	model.HookMap["gormDBExec"] = new(GormDBExec)
}

type GormDBExec struct {
}

func (h *GormDBExec) Hook() {
	db := &gorm.DB{}
	err := gohook.HookMethod(db, "Exec", Exec, ExecT)
	if err != nil {
		fmt.Println(err, "gormExec")
	} else {
		fmt.Println("gormExec")
	}
}

func (h *GormDBExec) UnHook() {
	db := &gorm.DB{}
	gohook.UnHookMethod(db, "Exec")
}
