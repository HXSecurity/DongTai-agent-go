package httpServeHTTP

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
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
	if err != nil {
		fmt.Println(err, "HttpServeHTTP")
	} else {
		fmt.Println("HttpServeHTTP")
	}
}

func (h *HttpServeHTTP) UnHook() {
	mux := &http.ServeMux{}
	gohook.UnHookMethod(mux, "ServeHTTP")
}
