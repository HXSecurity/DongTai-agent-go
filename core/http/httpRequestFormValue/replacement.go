package httpRequestFormValue

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"net/http"
)

func FormValue(req *http.Request, key string) string {
	r := FormValueT(req, key)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(key),
		Reqs:            utils.Collect(r),
		Source:          true,
		OriginClassName: "http.(*Request)",
		MethodName:      "FormValue",
		ClassName:       "http.(*Request)",
	})
	return r
}

func FormValueT(r *http.Request, key string) string {
	return key
}
