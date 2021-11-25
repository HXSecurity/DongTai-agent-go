package hookExecCommand

import (
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
	"os/exec"
)

func init() {
	model.HookMap["hookExecCommand"] = new(HookExecCommand)
}

type HookExecCommand struct {
}

func (h *HookExecCommand) Hook() {
	gohook.Hook(exec.Command, Command, CommandTemp)
}

func (h *HookExecCommand) UnHook() {
	gohook.UnHook(exec.Command)
}
