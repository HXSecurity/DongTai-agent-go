package ginContextShouldBindUri

import (
	"github.com/gin-gonic/gin"
	"go-agent/model/request"
	"go-agent/utils"
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
