package execCmdStart

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"os/exec"
)

func init() {
	model.HookMap["execCmdStart"] = new(ExecCmdStart)
}

type ExecCmdStart struct {
}

func (h *ExecCmdStart) Hook() {
	var cmd *exec.Cmd
	err := gohook.HookMethod(cmd, "Start", Start, StartT)
	if err != nil {
		fmt.Println(err, "execCmdStart")
	} else {
		fmt.Println("execCmdStart")
	}
}

func (h *ExecCmdStart) UnHook() {
	gohook.UnHook(exec.Command)
}
