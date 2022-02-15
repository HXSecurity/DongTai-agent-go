package httpClientDo

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"net/http"
	"reflect"
)

func Do(c *http.Client, req *http.Request) (*http.Response, error) {
	res, err := DoR(c, req)
	var u uintptr
	value := reflect.ValueOf(req)
	u = value.Pointer()
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(c, req),
		Reqs:            utils.Collect(res, err),
		NeedHook:        utils.Collect(u),
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
