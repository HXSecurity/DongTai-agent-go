package hook

import (
	"fmt"
	"go-agent/model"
)

//此处为定义了Hook的注册方法和卸载方法

func HookFunc(s string) {
	if model.HookMap[s] == nil {
		fmt.Println("尚未注册此HOOK")
	} else {
		model.HookMap[s].Hook()
	}
}

func UnHookFunc(s string) {
	if model.HookMap[s] == nil {
		fmt.Println("尚未注册此HOOK")
	} else {
		model.HookMap[s].UnHook()
	}
}
