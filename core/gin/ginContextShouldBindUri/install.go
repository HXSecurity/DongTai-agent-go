package ginContextShouldBindUri

import (
	"fmt"
	"github.com/brahma-adshonor/gohook"
	"github.com/gin-gonic/gin"
	"go-agent/model"
)

func init() {
	model.HookMap["ginContextShouldBindUri"] = new(GinContextShouldBindUri)
}

type GinContextShouldBindUri struct {
}

func (h *GinContextShouldBindUri) Hook() {
	context := &gin.Context{}
	err := gohook.HookMethod(context, "ShouldBindUri", ShouldBindUri, ShouldBindUriT)
	if err != nil {
		fmt.Println(err, "GinContextShouldBindUri")
	} else {
		fmt.Println("GinContextShouldBindUri")
	}
}

func (h *GinContextShouldBindUri) UnHook() {
	context := &gin.Context{}
	gohook.UnHookMethod(context, "ShouldBindUri")
}
