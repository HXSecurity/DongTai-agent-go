package ioReadAll

import (
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
	"io"
)

func init() {
	model.HookMap["ioReadAll"] = new(IoReadAll)
}

type IoReadAll struct {
}

func (h *IoReadAll) Hook() {
	gohook.Hook(io.ReadAll, ReadAll, ReadAllT)
}

func (h *IoReadAll) UnHook() {
	gohook.UnHook(io.ReadAll)
}
