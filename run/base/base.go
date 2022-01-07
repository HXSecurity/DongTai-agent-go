package base

import (
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/fmtSprintf"
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/jsonDecoderDecode"
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/jsonNewDecoder"
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/jsonUnmarshal"
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/runtimeConcatstrings"
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/sqlDBQuery"
	"github.com/HXSecurity/DongTai-agent-go/global"
	"github.com/HXSecurity/DongTai-agent-go/hook"
	"github.com/HXSecurity/DongTai-agent-go/service"
)

func init() {
	global.InitViper()
	service.AgentRegister()
	hook.HookFunc("sqlDBQuery")
	hook.HookFunc("fmtSprintf")
	hook.HookFunc("jsonUnmarshal")
	hook.HookFunc("jsonDecoderDecode")
	hook.HookFunc("jsonNewDecoder")
	hook.HookFunc("runtimeConcatstrings")

	////	此处为探针安装完成后的生命周期 注册 Agent 唤醒心跳等都可以从这里开始
	//req := request.AgentRegisterReq{
	//	Name:             "Mac OS X-localhost-v1.0.0-61862e3851934b9d96f34808e6354a5f",
	//	Language:         "GO",
	//	Version:          "v1.0.0",
	//	ProjectName:      "webgoat",
	//	Hostname:         "localhost",
	//	Network:          "{\"name\":\"feth3293\",\"ip\":\"172.30.181.207\"},{\"name\":\"en0\",\"ip\":\"192.168.2.161\"}",
	//	ContainerName:    "Apache Tomcat/9.0.37",
	//	ContainerVersion: "9.0.37.0",
	//	ServerAddr:       "",
	//	ServerPort:       "",
	//	ServerPath:       "/Users/x x x/workspace/",
	//	ServerEnv:        "JTdCJTIybmFtZSUyMiUzQSUyMDElMkMlMjAlMjJwYXRoJTIyJTNBJTIwJTIyL2Jpbi8lMjIlN0Q=",
	//	Pid:              "75495",
	//}
	//api.AgentRegister(req)
	//api.Profiles(request.HookRuleReq{Language: "GO"})
}

func main() {
	//go func() {
	//	fmt.Println(utils.CatGoroutineID())
	//	//获取当前开启的go协成的唯一ID 同一个协程内部完全相同
	//}()
	//go func() {
	//	fmt.Println(utils.CatGoroutineID())
	//	go func() {
	//		fmt.Println(utils.CatGoroutineID())
	//
	//	}()
	//	//获取当前开启的go协成的唯一ID 同一个协程内部完全相同
	//
	//}()
	//go func() {
	//	fmt.Println(utils.CatGoroutineID())
	//
	//	//获取当前开启的go协成的唯一ID 同一个协程内部完全相同
	//
	//}()
	//go func() {
	//	fmt.Println(utils.CatGoroutineID())
	//
	//	//获取当前开启的go协成的唯一ID 同一个协程内部完全相同
	//}()
	//fmt.Println(utils.CatGoroutineID())
	//// 上述四个为测试携程 最后为主携程
	//for {
	//
	//}
	//// 此处模拟的接口调用后 对hook map中的声明过的hook进行Hook 和 Unhook 用户不会执行到main 但是引入包必然会触发init 的生命周期
	//hook.HookFunc("http")
	//hook.HookFunc("exec")
	//
	//c:= new(http.Client)
	//c.Get("https://www.baidu.com")
	//此处的原生http已被自己开发的http get接口方法捕捉
	//exec.Command("ls")
	//此处的原生Command已被自己开发的Command接口方法捕捉
	//// 注销方法
	//
	//hook.UnHookFunc("http")
	//卸载捕获
	//hook.UnHookFunc("exec")
	// 卸载捕获
	//
	//
	//c.Get("https://www.baidu123.com")
	// 已经恢复到原生
	//
	//exec.Command("clear")
}
