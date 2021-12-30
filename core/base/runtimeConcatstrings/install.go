package runtimeConcatstrings

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
)

func init() {
	model.HookMap["runtimeConcatstrings"] = new(RuntimeConcatstrings)
}

type RuntimeConcatstrings struct {
}

func (h *RuntimeConcatstrings) Hook() {
	err := gohook.Hook(concatstrings, concatstringsR, concatstringsT)
	if err != nil {
		fmt.Println(err, "RuntimeConcatstrings")
	} else {
		fmt.Println("RuntimeConcatstrings")
	}
}

func (h *RuntimeConcatstrings) UnHook() {
	gohook.UnHook(concatstrings)
}
