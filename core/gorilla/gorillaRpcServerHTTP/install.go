package gorillaRpcServerHTTP

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"github.com/gorilla/rpc/v2"
)

func init() {
	model.HookMap["gorillaRpcServerHTTP"] = new(GorillaRpcServerHTTP)
}

type GorillaRpcServerHTTP struct {
}

func (h *GorillaRpcServerHTTP) Hook() {
	mux := &rpc.Server{}
	err := gohook.HookMethod(mux, "ServeHTTP", MyServer, MyServerTemp)
	if err != nil {
		fmt.Println(err, "GorillaRpcServerHTTP")
	} else {
		fmt.Println("GorillaRpcServerHTTP")
	}
}

func (h *GorillaRpcServerHTTP) UnHook() {
	mux := &rpc.Server{}
	gohook.UnHookMethod(mux, "ServeHTTP")
}
