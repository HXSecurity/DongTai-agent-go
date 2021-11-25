package hookSprintf

import (
	"fmt"
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
)

func init() {
	model.HookMap["hookSprintf"] = new(HookSprintf)
}

type HookSprintf struct {
}

func (h *HookSprintf) Hook() {
	gohook.Hook(fmt.Sprintf, Sprintf, SprintfT)
}

func (h *HookSprintf) UnHook() {
	gohook.UnHook(fmt.Sprintf)
}
