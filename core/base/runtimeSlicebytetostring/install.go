package runtimeSlicebytetostring

import (
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
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
