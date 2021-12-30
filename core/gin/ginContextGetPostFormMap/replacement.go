package ginContextGetPostFormMap

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"github.com/gin-gonic/gin"
)

func GetPostFormMap(c *gin.Context, key string) (map[string]string, bool) {
	reMap, flag := GetPostFormMapT(c, key)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(key),
		Reqs:            utils.Collect(reMap),
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
