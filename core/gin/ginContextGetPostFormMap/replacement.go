package ginContextGetPostFormMap

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/gin-gonic/gin"
)

func GetPostFormMap(c *gin.Context, key string) (map[string]string, bool) {
	reMap, flag := GetPostFormMapT(c, key)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(key),
		Reqs:            request.Collect(reMap),
		Source:          true,
		OriginClassName: "gin.(*Context)",
		MethodName:      "GetPostFormMap",
		ClassName:       "gin.(*Context)",
	})
	return reMap, flag
}

func GetPostFormMapT(c *gin.Context, key string) (map[string]string, bool) {
	return make(map[string]string), false
}
