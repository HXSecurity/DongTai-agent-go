package hookJoin

import (
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
	"strings"
)

func init() {
	model.HookMap["hookJoin"] = new(HookJoin)
}

type HookJoin struct {
}

func (h *HookJoin) Hook() {
	gohook.Hook(strings.Join, Join, JoinT)
}

func (h *HookJoin) UnHook() {
	gohook.UnHook(strings.Join)
}
