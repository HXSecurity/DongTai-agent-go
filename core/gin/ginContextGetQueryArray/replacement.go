package ginContextGetQueryArray

import (
	"github.com/gin-gonic/gin"
	"go-agent/model/request"
	"go-agent/utils"
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
