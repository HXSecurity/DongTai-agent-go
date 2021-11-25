package httpRouter

import (
	"fmt"
	"github.com/brahma-adshonor/gohook"
	"github.com/julienschmidt/httprouter"
	"go-agent/model"
)

func init() {
	model.HookMap["httpRouter"] = new(HttpRouter)
}

type HttpRouter struct {
}

func (h *HttpRouter) Hook() {
	mux := &httprouter.Router{}
	err := gohook.HookMethod(mux, "ServeHTTP", MyHttpRouterServer, MyHttpRouterServerTemp)
	fmt.Println(err)
}

func (h *HttpRouter) UnHook() {
	mux := &httprouter.Router{}
	gohook.UnHookMethod(mux, "ServeHTTP")
}
