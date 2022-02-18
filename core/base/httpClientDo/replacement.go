package httpClientDo

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"net/http"
	"reflect"
)

func Do(c *http.Client, req *http.Request) (*http.Response, error) {
	res, err := DoR(c, req)
	var u uintptr
	value := reflect.ValueOf(req)
	u = value.Pointer()
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(c, req),
		Reqs:            request.Collect(res, err),
		NeedHook:        request.Collect(u),
		Source:          false,
		OriginClassName: "http.(*Client)",
		MethodName:      "Do",
		ClassName:       "http.(*Client)",
	})
	return res, err
}

func DoR(c *http.Client, req *http.Request) (*http.Response, error) {
	return nil, nil
}
