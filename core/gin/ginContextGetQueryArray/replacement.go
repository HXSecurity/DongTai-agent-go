package ginContextGetQueryArray

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/gin-gonic/gin"
)

func GetQueryArray(c *gin.Context, key string) ([]string, bool) {
	reArray, flag := GetQueryArrayT(c, key)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(key),
		Reqs:            request.Collect(reArray),
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
