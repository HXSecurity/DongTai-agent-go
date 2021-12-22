package ginEngineServerHTTP

import (
	"fmt"
	"github.com/brahma-adshonor/gohook"
	"github.com/gin-gonic/gin"
	"go-agent/model"
)

func init() {
	model.HookMap["ginEngineServerHTTP"] = new(GinEngineServerHTTP)
}

type GinEngineServerHTTP struct {
}

func (h *GinEngineServerHTTP) Hook() {
	mux := &gin.Engine{}
	err := gohook.HookMethod(mux, "ServeHTTP", MyServer, MyServerTemp)
	if err != nil {
		fmt.Println(err, "GinEngineServerHTTP")
	} else {
		fmt.Println("GinEngineServerHTTP")
	}
}

func (h *GinEngineServerHTTP) UnHook() {
	mux := &gin.Engine{}
	gohook.UnHookMethod(mux, "ServeHTTP")
}
