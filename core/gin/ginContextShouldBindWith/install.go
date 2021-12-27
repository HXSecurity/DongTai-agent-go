package ginContextShouldBindWith

import (
	"fmt"
	"github.com/brahma-adshonor/gohook"
	"github.com/gin-gonic/gin"
	"go-agent/model"
)

func init() {
	model.HookMap["ginContextShouldBindWith"] = new(GinContextShouldBindWith)
}

type GinContextShouldBindWith struct {
}

func (h *GinContextShouldBindWith) Hook() {
	context := &gin.Context{}
	err := gohook.HookMethod(context, "ShouldBindWith", ShouldBindWith, ShouldBindWithT)
	if err != nil {
		fmt.Println(err, "GinContextShouldBindWith")
	} else {
		fmt.Println("GinContextShouldBindWith")
	}
}

func (h *GinContextShouldBindWith) UnHook() {
	context := &gin.Context{}
	gohook.UnHookMethod(context, "ShouldBindWith")
}
