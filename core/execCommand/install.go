package execCommand

import (
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
	"os/exec"
)

func init() {
	model.HookMap["execCommand"] = new(ExecCommand)
}

type ExecCommand struct {
}

func (h *ExecCommand) Hook() {
	gohook.Hook(exec.Command, Command, CommandTemp)
}

func (h *ExecCommand) UnHook() {
	gohook.UnHook(exec.Command)
}
