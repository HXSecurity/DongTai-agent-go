package runtimeSlicebytetostring

import (
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
)

func init() {
	model.HookMap["runtimeSlicebytetostring"] = new(RuntimeSlicebytetostring)
}

type RuntimeSlicebytetostring struct {
}

func (h *RuntimeSlicebytetostring) Hook() {
	gohook.Hook(slicebytetostring, slicebytetostringR, slicebytetostringT)
}

func (h *RuntimeSlicebytetostring) UnHook() {
	gohook.UnHook(slicebytetostring)
}
