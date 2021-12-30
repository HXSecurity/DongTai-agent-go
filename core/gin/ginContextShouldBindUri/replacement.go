package ginContextShouldBindUri

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"github.com/gin-gonic/gin"
)

func ShouldBindUri(c *gin.Context, obj interface{}) error {
	err := ShouldBindUriT(c, obj)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(),
		Reqs:            utils.Collect(obj),
		Source:          true,
		OriginClassName: "gin.(*Context)",
		MethodName:      "ShouldBindUri",
		ClassName:       "gin.(*Context)",
	})
	return err
}

func ShouldBindUriT(c *gin.Context, obj interface{}) error {
	return nil
}
