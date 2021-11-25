package stringsJoin

import (
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
	"strings"
)

func init() {
	model.HookMap["stringsJoin"] = new(StringsJoin)
}

type StringsJoin struct {
}

func (h *StringsJoin) Hook() {
	gohook.Hook(strings.Join, Join, JoinT)
}

func (h *StringsJoin) UnHook() {
	gohook.UnHook(strings.Join)
}
