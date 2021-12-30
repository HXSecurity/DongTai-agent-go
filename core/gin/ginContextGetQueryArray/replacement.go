package ginContextGetQueryArray

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"github.com/gin-gonic/gin"
)

func GetQueryArray(c *gin.Context, key string) ([]string, bool) {
	reArray, flag := GetQueryArrayT(c, key)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(key),
		Reqs:            utils.Collect(reArray),
		Source:          true,
		OriginClassName: "gin.(*Context)",
		MethodName:      "GetQueryArray",
		ClassName:       "gin.(*Context)",
	})
	return reArray, flag
}

func GetQueryArrayT(c *gin.Context, key string) ([]string, bool) {
	return []string{}, false
}
