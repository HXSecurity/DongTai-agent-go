package ginContextGetPostFormArray

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"github.com/gin-gonic/gin"
)

func init() {
	model.HookMap["ginContextGetPostFormArray"] = new(GinContextGetPostFormArray)
}

type GinContextGetPostFormArray struct {
}

func (h *GinContextGetPostFormArray) Hook() {
	context := &gin.Context{}
	err := gohook.HookMethod(context, "GetPostFormArray", GetPostFormArray, GetPostFormArrayT)
	if err != nil {
		fmt.Println(err, "GinContextGetPostFormArray")
	} else {
		fmt.Println("GinContextGetPostFormArray")
	}
}

func (h *GinContextGetPostFormArray) UnHook() {
	context := &gin.Context{}
	gohook.UnHookMethod(context, "GetPostFormArray")
}
