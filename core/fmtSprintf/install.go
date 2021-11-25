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
	gohook.Hook(fmt.Sprintf, Sprintf, SprintfT)
}

func (h *FmtSprintf) UnHook() {
	gohook.UnHook(fmt.Sprintf)
}
