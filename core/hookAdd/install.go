package hookAdd

import (
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
	"os/exec"
)

// 模拟了对exec 包 Command 的hook  此处不规范 仅为展示效果

func init() {
	model.HookMap["add"] = new(HookAdd)
}

type HookAdd struct {
}

func (h *HookAdd) Hook() {
	gohook.Hook(concatstrings, concatstringsR, concatstringsT)
}

func (h *HookAdd) UnHook() {
	gohook.UnHook(exec.Command)
}
