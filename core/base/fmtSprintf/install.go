package fmtSprintf

import (
	"fmt"
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
)

func init() {
	model.HookMap["fmtSprintf"] = new(FmtSprintf)
}

type FmtSprintf struct {
}

func (h *FmtSprintf) Hook() {
	err := gohook.Hook(fmt.Sprintf, Sprintf, SprintfT)
	if err != nil {
		fmt.Println(err, "Sprintf")
	} else {
		fmt.Println("Sprintf")
	}
}

func (h *FmtSprintf) UnHook() {
	gohook.UnHook(fmt.Sprintf)
}
