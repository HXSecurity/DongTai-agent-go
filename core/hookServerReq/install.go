package hookServerReq

import (
	"fmt"
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
	"net/http"
)

func init() {
	model.HookMap["serverReq"] = new(HookServerReq)
}

type HookServerReq struct {
}

func (h *HookServerReq) Hook() {
	mux := &http.ServeMux{}
	err := gohook.HookMethod(mux, "ServeHTTP", MyServerReq, MyServerTempReq)
	fmt.Println(err)
}

func (h *HookServerReq) UnHook() {
	mux := &http.ServeMux{}
	gohook.UnHookMethod(mux, "ServeHTTP")
}
