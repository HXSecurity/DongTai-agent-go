package hookConcatstrings

import (
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
)

func init() {
	model.HookMap["hookConcatstrings"] = new(HookConcatstrings)
}

type HookConcatstrings struct {
}

func (h *HookConcatstrings) Hook() {
	gohook.Hook(concatstrings, concatstringsR, concatstringsT)
}

func (h *HookConcatstrings) UnHook() {
	gohook.UnHook(concatstrings)
}
