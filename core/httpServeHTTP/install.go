package httpServeHTTP

import (
	"fmt"
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
	"net/http"
)

func init() {
	model.HookMap["httpServeHTTP"] = new(HttpServeHTTP)
}

type HttpServeHTTP struct {
}

func (h *HttpServeHTTP) Hook() {
	mux := &http.ServeMux{}
	err := gohook.HookMethod(mux, "ServeHTTP", MyServer, MyServerTemp)
	fmt.Println(err)
}

func (h *HttpServeHTTP) UnHook() {
	mux := &http.ServeMux{}
	gohook.UnHookMethod(mux, "ServeHTTP")
}
