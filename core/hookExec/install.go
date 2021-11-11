package hookExec

import (
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
	"os/exec"
)

// 模拟了对exec 包 Command 的hook  此处不规范 仅为展示效果


func init(){
	model.HookMap["exec"] = new(HookExec)
}

type HookExec struct {

}

func (h *HookExec)Hook()  {
	gohook.Hook(exec.Command, Command, CommandTemp)
}


func (h *HookExec)UnHook()  {
	gohook.UnHook(exec.Command)
}