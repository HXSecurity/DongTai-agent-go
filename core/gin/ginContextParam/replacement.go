package ginContextParam

import (
	"github.com/gin-gonic/gin"
	"go-agent/model/request"
	"go-agent/utils"
)

func Param(c *gin.Context, key string) string {
	str := ParamT(c, key)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(key),
		Reqs:            utils.Collect(str),
		Source:          true,
		OriginClassName: "gin.(*Context)",
		MethodName:      "Param",
		ClassName:       "gin.(*Context)",
	})
	return str
}

func ParamT(c *gin.Context, key string) string {
	return ""
}
