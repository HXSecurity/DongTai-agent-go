package runtimeMapassign_faststr

import (
	"fmt"
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
)

func init() {
	model.HookMap["runtimeConcatstrings"] = new(RuntimeMapassign_faststr)
}

type RuntimeMapassign_faststr struct {
}

func (h *RuntimeMapassign_faststr) Hook() {
	err := gohook.Hook(mapassign_faststr, mapassign_faststrR, mapassign_faststrT)
	if err != nil {
		fmt.Println(err, "RuntimeMapassign_faststr")
	} else {
		fmt.Println("RuntimeMapassign_faststr")
	}
}

func (h *RuntimeMapassign_faststr) UnHook() {
	gohook.UnHook(mapassign_faststr)
}
