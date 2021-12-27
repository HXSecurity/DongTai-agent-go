package ginContextParam

import (
	"fmt"
	"github.com/brahma-adshonor/gohook"
	"github.com/gin-gonic/gin"
	"go-agent/model"
)

func init() {
	model.HookMap["ginContextParam"] = new(GinContextParam)
}

type GinContextParam struct {
}

func (h *GinContextParam) Hook() {
	context := &gin.Context{}
	err := gohook.HookMethod(context, "Param", Param, ParamT)
	if err != nil {
		fmt.Println(err, "GinContextParam")
	} else {
		fmt.Println("GinContextParam")
	}
}

func (h *GinContextParam) UnHook() {
	context := &gin.Context{}
	gohook.UnHookMethod(context, "Param")
}
