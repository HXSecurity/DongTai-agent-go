package runtimesSringtoslicebyte

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
)

func init() {
	model.HookMap["runtimesSringtoslicebyte"] = new(RuntimesSringtoslicebyte)
}

type RuntimesSringtoslicebyte struct {
}

func (h *RuntimesSringtoslicebyte) Hook() {
	err := gohook.Hook(stringtoslicebyte, stringtoslicebyteR, stringtoslicebyteT)
	if err != nil {
		fmt.Println("RuntimesSringtoslicebyte:", err)
	} else {
		fmt.Println("RuntimesSringtoslicebyte")
	}
}

func (h *RuntimesSringtoslicebyte) UnHook() {
	err := gohook.UnHook(stringtoslicebyte)
	if err != nil {
		fmt.Println("RuntimesSringtoslicebyte:", err)
	} else {
		fmt.Println("RuntimesSringtoslicebyte")
	}
}
