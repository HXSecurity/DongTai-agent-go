package hook

import (
	"fmt"
	"go-agent/model"
)

//此处为定义了Hook的注册方法和卸载方法

func HookFunc(string string){
	if(model.HookMap[string] == nil){
		fmt.Println("尚未注册此HOOK")
	}else{
		model.HookMap[string].Hook()
	}
}



func UnHookFunc(string string){
	if(model.HookMap[string] == nil){
		fmt.Println("尚未注册此HOOK")
	}else{
		model.HookMap[string].UnHook()
	}
}
