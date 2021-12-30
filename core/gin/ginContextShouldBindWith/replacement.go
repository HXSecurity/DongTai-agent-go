package ginContextShouldBindWith

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ShouldBindWith(c *gin.Context, obj interface{}, b binding.Binding) error {
	err := ShouldBindWithT(c, obj, b)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(),
		Reqs:            utils.Collect(obj),
		Source:          true,
		OriginClassName: "gin.(*Context)",
		MethodName:      "ShouldBindWith",
		ClassName:       "gin.(*Context)",
	})
	return err
}

func ShouldBindWithT(c *gin.Context, obj interface{}, b binding.Binding) error {
	return nil
}
