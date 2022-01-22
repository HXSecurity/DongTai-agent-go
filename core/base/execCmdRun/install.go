package execCmdRun

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"os/exec"
)

func init() {
	model.HookMap["execCmdRun"] = new(ExecCmdRun)
}

type ExecCmdRun struct {
}

func (h *ExecCmdRun) Hook() {
	var cmd *exec.Cmd
	err := gohook.HookMethod(cmd, "Run", Run, RunT)
	if err != nil {
		fmt.Println(err, "execCmdRun")
	} else {
		fmt.Println("execCmdRun")
	}
}

func (h *ExecCmdRun) UnHook() {
	gohook.UnHook(exec.Command)
}
