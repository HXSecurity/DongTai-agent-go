package model

//定义的方法接口 后续更多功能 最好是嵌套接口
type HookFunc interface {
	Hook()
	UnHook()
}

// 接口扩展时期示例  不影响原有功能
type More interface {
	HookFunc
	More()
}

//实现一个方法调用的接口集合 用于对各类hook的统一map管理 动态挂载卸载
var HookMap = make(map[string]HookFunc)

type HookStruct interface {
	GetHook() []string
	HookAll()
	UnHookAll()
}
