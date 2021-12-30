package gormDBOrder

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"gorm.io/gorm"
)

func init() {
	model.HookMap["gormDBOrder"] = new(GormDBOrder)
}

type GormDBOrder struct {
}

func (h *GormDBOrder) Hook() {
	db := &gorm.DB{}
	err := gohook.HookMethod(db, "Order", Order, OrderT)
	if err != nil {
		fmt.Println(err, "gormOrder")
	} else {
		fmt.Println("gormOrder")
	}
}

func (h *GormDBOrder) UnHook() {
	db := &gorm.DB{}
	gohook.UnHookMethod(db, "Order")
}
