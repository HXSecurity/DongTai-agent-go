package ginContextGetPostFormArray

import (
	"github.com/gin-gonic/gin"
	"go-agent/model/request"
	"go-agent/utils"
)

func GetPostFormArray(c *gin.Context, key string) ([]string, bool) {
	reArray, flag := GetPostFormArrayT(c, key)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(key),
		Reqs:            utils.Collect(reArray),
		Source:          true,
		OriginClassName: "gin.(*Context)",
		MethodName:      "GetPostFormArray",
		ClassName:       "gin.(*Context)",
	})
	return reArray, flag
}

func GetPostFormArrayT(c *gin.Context, key string) ([]string, bool) {
	return []string{}, false
}
