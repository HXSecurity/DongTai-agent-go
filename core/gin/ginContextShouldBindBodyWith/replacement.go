package ginContextShouldBindBodyWith

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
