package runtimeConcatstrings

import (
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
)

func init() {
	model.HookMap["runtimeConcatstrings"] = new(RuntimeConcatstrings)
}

type RuntimeConcatstrings struct {
}

func (h *RuntimeConcatstrings) Hook() {
	gohook.Hook(concatstrings, concatstringsR, concatstringsT)
}

func (h *RuntimeConcatstrings) UnHook() {
	gohook.UnHook(concatstrings)
}
