package httpNewRequest

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"io"
	"net/http"
	"reflect"
)

func NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := NewRequestR(method, url, body)
	var u uintptr
	value := reflect.ValueOf(req)
	u = value.Pointer()
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(method, url, body),
		Reqs:            request.Collect(req, err),
		NeedCatch:       request.Collect(u),
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
