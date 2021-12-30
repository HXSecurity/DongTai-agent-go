package ginContextParam

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"github.com/gin-gonic/gin"
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
