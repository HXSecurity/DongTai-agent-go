package ginContextShouldBindBodyWith

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-agent/model/request"
	"go-agent/utils"
)

func ShouldBindBodyWith(c *gin.Context, obj interface{}, b binding.BindingBody) error {
	err := ShouldBindBodyWithT(c, obj, b)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(),
		Reqs:            utils.Collect(obj),
		Source:          true,
		OriginClassName: "gin.(*Context)",
		MethodName:      "ShouldBindBodyWith",
		ClassName:       "gin.(*Context)",
	})
	return err
}

func ShouldBindBodyWithT(c *gin.Context, obj interface{}, b binding.BindingBody) error {
	return nil
}
