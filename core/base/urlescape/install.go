package urlescape

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
)

func init() {
	model.HookMap["urlescape"] = new(Urlescape)
}

type Urlescape struct {
}

func (h *Urlescape) Hook() {
	err := gohook.Hook(escape, escapeR, escapeT)
	if err != nil {
		fmt.Println("Urlescape:", err)
	} else {
		fmt.Println("Urlescape")
	}
}

func (h *Urlescape) UnHook() {
	err := gohook.UnHook(escape)
	if err != nil {
		fmt.Println("Urlescape:", err)
	} else {
		fmt.Println("Urlescape")
	}
}
