package hookServer

import (
	"fmt"
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
	"net/http"
)

func init() {
	model.HookMap["hookServer"] = new(HookServer)
}

type HookServer struct {
}

func (h *HookServer) Hook() {
	mux := &http.ServeMux{}
	err := gohook.HookMethod(mux, "ServeHTTP", MyServerReq, MyServerTempReq)
	fmt.Println(err)
}

func (h *HookServer) UnHook() {
	mux := &http.ServeMux{}
	gohook.UnHookMethod(mux, "ServeHTTP")
}
