package httpRequestFormValue

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"net/http"
)

func FormValue(req *http.Request, key string) string {
	r := FormValueT(req, key)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(key),
		Reqs:            request.Collect(r),
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
