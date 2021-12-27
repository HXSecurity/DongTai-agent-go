package ginContextGetQueryMap

import (
	"fmt"
	"github.com/brahma-adshonor/gohook"
	"github.com/gin-gonic/gin"
	"go-agent/model"
)

func init() {
	model.HookMap["ginContextGetQueryMap"] = new(GinContextGetQueryMap)
}

type GinContextGetQueryMap struct {
}

func (h *GinContextGetQueryMap) Hook() {
	context := &gin.Context{}
	err := gohook.HookMethod(context, "GetQueryMap", GetQueryMap, GetQueryMapT)
	if err != nil {
		fmt.Println(err, "GinContextGetQueryMap")
	} else {
		fmt.Println("GinContextGetQueryMap")
	}
}

func (h *GinContextGetQueryMap) UnHook() {
	context := &gin.Context{}
	gohook.UnHookMethod(context, "GetQueryMap")
}
