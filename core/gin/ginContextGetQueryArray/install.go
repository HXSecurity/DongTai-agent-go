package ginContextGetQueryArray

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"github.com/gin-gonic/gin"
)

func init() {
	model.HookMap["ginContextGetQueryArray"] = new(GinContextGetQueryArray)
}

type GinContextGetQueryArray struct {
}

func (h *GinContextGetQueryArray) Hook() {
	context := &gin.Context{}
	err := gohook.HookMethod(context, "GetQueryArray", GetQueryArray, GetQueryArrayT)
	if err != nil {
		fmt.Println(err, "GinContextGetQueryArray")
	} else {
		fmt.Println("GinContextGetQueryArray")
	}
}

func (h *GinContextGetQueryArray) UnHook() {
	context := &gin.Context{}
	gohook.UnHookMethod(context, "GetQueryArray")
}
