package main

import (
	"go-agent/hook"
	_ "go-agent/core/hookExec"
	_ "go-agent/core/hookhttp"
	"net/http"
	"os/exec"
)



func init() {
//	此处为探针安装完成后的生命周期 注册 Agent 唤醒心跳等都可以从这里开始
}

func main() {
	// 此处模拟的接口调用后 对hook map中的声明过的hook进行Hook 和 Unhook 用户不会执行到main 但是引入包必然会触发init 的生命周期
	hook.HookFunc("http")
	hook.HookFunc("exec")

	c:= new(http.Client)
	c.Get("https://www.baidu.com")

	exec.Command("ls")
	// 注销方法

	hook.UnHookFunc("http")
	hook.UnHookFunc("exec")


	c.Get("https://www.baidu123.com")

	exec.Command("clear")
}
