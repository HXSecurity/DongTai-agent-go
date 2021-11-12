package main

import (
	"go-agent/api"
	"go-agent/global"
	"go-agent/model/request"
	"go-agent/utils"
)

func init() {
	global.InitViper()
	//	此处为探针安装完成后的生命周期 注册 Agent 唤醒心跳等都可以从这里开始
	req := request.AgentRegisterReq{
		Name:             "Mac OS X-localhost-v1.0.0-61862e3851934b9d96f34808e6354a5f",
		Language:         "GO",
		Version:          "v1.0.0",
		ProjectName:      "webgoat",
		Hostname:         "localhost",
		Network:          "{\"name\":\"feth3293\",\"ip\":\"172.30.181.207\"},{\"name\":\"en0\",\"ip\":\"192.168.2.161\"}",
		ContainerName:    "Apache Tomcat/9.0.37",
		ContainerVersion: "9.0.37.0",
		ServerAddr:       "",
		ServerPort:       "",
		ServerPath:       "/Users/x x x/workspace/",
		ServerEnv:        "JTdCJTIybmFtZSUyMiUzQSUyMDElMkMlMjAlMjJwYXRoJTIyJTNBJTIwJTIyL2Jpbi8lMjIlN0Q=",
		Pid:              "75495",
	}
	api.AgentRegister(req)
	api.Profiles(request.HookRuleReq{Language: "GO"})
}

func main() {
	go func() {
		utils.CatGoroutineID()
	}()
	go func() {
		utils.CatGoroutineID()

	}()
	go func() {
		utils.CatGoroutineID()
		utils.CatGoroutineID()
		utils.CatGoroutineID()

	}()
	go func() {
		utils.CatGoroutineID()

	}()
	utils.CatGoroutineID()
	utils.CatGoroutineID()
	utils.CatGoroutineID()
	for {

	}
	//// 此处模拟的接口调用后 对hook map中的声明过的hook进行Hook 和 Unhook 用户不会执行到main 但是引入包必然会触发init 的生命周期
	//hook.HookFunc("http")
	//hook.HookFunc("exec")
	//
	//c:= new(http.Client)
	//c.Get("https://www.baidu.com")
	//
	//exec.Command("ls")
	//// 注销方法
	//
	//hook.UnHookFunc("http")
	//hook.UnHookFunc("exec")
	//
	//
	//c.Get("https://www.baidu123.com")
	//
	//exec.Command("clear")
}
