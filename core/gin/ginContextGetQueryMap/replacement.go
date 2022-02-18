package ginContextGetQueryMap

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"github.com/gin-gonic/gin"
)

func GetQueryMap(c *gin.Context, key string) (map[string]string, bool) {
	reMap, flag := GetQueryMapT(c, key)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(key),
		Reqs:            request.Collect(reMap),
		Source:          true,
		OriginClassName: "gin.(*Context)",
		MethodName:      "GetQueryMap",
		ClassName:       "gin.(*Context)",
	})
	return reMap, flag
}

func GetQueryMapT(c *gin.Context, key string) (map[string]string, bool) {
	return make(map[string]string), false
}
