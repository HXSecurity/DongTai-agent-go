package stringsRepeat

import (
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
	"strings"
)

func init() {
	model.HookMap["stringsRepeat"] = new(StringsRepeat)
}

type StringsRepeat struct {
}

func (h *StringsRepeat) Hook() {
	gohook.Hook(strings.Repeat, Repeat, RepeatR)
}

func (h *StringsRepeat) UnHook() {
	gohook.UnHook(strings.Repeat)
}
