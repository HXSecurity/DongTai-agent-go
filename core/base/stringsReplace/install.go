package stringsReplace

import (
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"strings"
)

func init() {
	model.HookMap["stringsReplace"] = new(StringsReplace)
}

type StringsReplace struct {
}

func (h *StringsReplace) Hook() {
	gohook.Hook(strings.Replace, Replace, ReplaceR)
}

func (h *StringsReplace) UnHook() {
	gohook.UnHook(strings.Replace)
}
