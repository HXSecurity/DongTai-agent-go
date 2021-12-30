package ginContextGetPostFormMap

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"github.com/gin-gonic/gin"
)

func init() {
	model.HookMap["ginContextGetPostFormMap"] = new(GinContextGetPostFormMap)
}

type GinContextGetPostFormMap struct {
}

func (h *GinContextGetPostFormMap) Hook() {
	context := &gin.Context{}
	err := gohook.HookMethod(context, "GetPostFormMap", GetPostFormMap, GetPostFormMapT)
	if err != nil {
		fmt.Println(err, "GinContextGetPostFormMap")
	} else {
		fmt.Println("GinContextGetPostFormMap")
	}
}

func (h *GinContextGetPostFormMap) UnHook() {
	context := &gin.Context{}
	gohook.UnHookMethod(context, "GetPostFormMap")
}
