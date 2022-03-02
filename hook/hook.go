package hook

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
)

//此处为定义了Hook的注册方法和卸载方法

func HookFunc(s string) {
	if model.HookMap[s] == nil {
		fmt.Println("尚未注册此HOOK:", s)
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

func Hook(funcs []string) {
	for i := range funcs {
		HookFunc(funcs[i])
	}
}

func UnHook(funcs []string) {
	for i := range funcs {
		UnHookFunc(funcs[i])
	}
}

func HookAll(h ...model.HookStruct) {
	for i := range h {
		h[i].HookAll()
	}
}

func UnHookAll(h ...model.HookStruct) {
	for i := range h {
		h[i].UnHookAll()
	}
}

func RunAllHook() {
	HookAll(
		new(Base),
		new(Gin),
		new(Gorilla),
		new(Gorm),
		new(Http),
		new(HttpRouter),
	)
}

func StopAllHook() {
	UnHookAll(
		new(Base),
		new(Gin),
		new(Gorilla),
		new(Gorm),
		new(Http),
		new(HttpRouter),
	)
}
