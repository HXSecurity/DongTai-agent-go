package main

import (
	"fmt"
	_ "go-agent/core/httpServeHTTP"
	_ "go-agent/core/runtimeConcatstrings"
	"go-agent/global"
	"go-agent/hook"
	"go-agent/service"
	"io"
	"net/http"
)

func init() {
	global.InitViper()
	service.AgentRegister()
}

func doRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	//fmt.Println("path", r.URL.Path)
	//fmt.Println("scheme", r.URL.Scheme)
	//for k, v := range r.Form {
	//    fmt.Println("key:", k)
	//    fmt.Println("value:", strings.Join(v, ""))
	//}
	w.Header().Set("test", "test")
	fmt.Fprintf(w, "service start...") //这个写入到w的是输出到客户端的 也可以用下面的 io.WriteString对象
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	//query := r.URL.Query()
	var uid string // 初始化定义变量
	if r.Method == "GET" {
		uid = r.FormValue("uid")
	} else if r.Method == "POST" {
		uid = r.PostFormValue("uid")
	}
	io.WriteString(w, "uid = "+uid)
}

func main() {
	hook.HookFunc("httpServeHTTP")
	hook.HookFunc("runtimeConcatstrings")
	a := "2" + "3"
	fmt.Println(a)
	//service.PingPang()
	http.HandleFunc("/test", doRequest) //   设置访问路由
	_ = http.ListenAndServe(":9090", nil)

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
	// 卸载捕获
	//hook.UnHookFunc("exec")
	// 卸载捕获
	//
	//
	//c.Get("https://www.baidu123.com")
	// 已经恢复到原生
	//
	//exec.Command("clear")
}
