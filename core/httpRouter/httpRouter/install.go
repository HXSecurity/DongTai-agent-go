package httpRouter

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"github.com/julienschmidt/httprouter"
)

func init() {
	model.HookMap["httpRouter"] = new(HttpRouter)
}

type HttpRouter struct {
}

func (h *HttpRouter) Hook() {
	mux := &httprouter.Router{}
	err := gohook.HookMethod(mux, "ServeHTTP", MyHttpRouterServer, MyHttpRouterServerTemp)
	if err != nil {
		fmt.Println(err, "HttpRouter")
	} else {
		fmt.Println("HttpRouter")
	}
}

func (h *HttpRouter) UnHook() {
	mux := &httprouter.Router{}
	gohook.UnHookMethod(mux, "ServeHTTP")
}
