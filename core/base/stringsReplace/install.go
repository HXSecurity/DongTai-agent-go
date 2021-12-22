package stringsReplace

import (
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
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
