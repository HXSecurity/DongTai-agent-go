package stringsRepeat

import (
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
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
