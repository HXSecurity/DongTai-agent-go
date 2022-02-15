package httpNewRequest

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"io"
	"net/http"
	"reflect"
)

func NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := NewRequestR(method, url, body)
	var u uintptr
	value := reflect.ValueOf(req)
	u = value.Pointer()
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(method, url, body),
		Reqs:            utils.Collect(req, err),
		NeedCatch:       utils.Collect(u),
		Source:          false,
		OriginClassName: "http",
		MethodName:      "NewRequest",
		ClassName:       "http",
	})
	return req, err
}

func NewRequestR(method, url string, body io.Reader) (*http.Request, error) {
	return nil, nil
}
