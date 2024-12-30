package muxServeHTTP

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"github.com/go-mux/mux"
)

func init() {
	model.HookMap["muxServeHTTP"] = new(MuxServeHTTP)
}

type MuxServeHTTP struct {
}

func (h *MuxServeHTTP) Hook() {
	mx := &mux.Router{}
	err := gohook.HookMethod(mx, "ServeHTTP", MyServer, MyServerTemp)
	if err != nil {
		fmt.Println(err, "HttpServeHTTP")
	} else {
		fmt.Println("HttpServeHTTP")
	}
}

func (h *MuxServeHTTP) UnHook() {
	mx := &mux.Router{}
	gohook.UnHookMethod(mx, "ServeHTTP")
}
