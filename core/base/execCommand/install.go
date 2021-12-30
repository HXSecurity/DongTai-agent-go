package execCommand

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"os/exec"
)

func init() {
	model.HookMap["execCommand"] = new(ExecCommand)
}

type ExecCommand struct {
}

func (h *ExecCommand) Hook() {
	err := gohook.Hook(exec.Command, Command, CommandTemp)
	if err != nil {
		fmt.Println(err, "ExecCommand")
	} else {
		fmt.Println("ExecCommand")
	}
}

func (h *ExecCommand) UnHook() {
	gohook.UnHook(exec.Command)
}
