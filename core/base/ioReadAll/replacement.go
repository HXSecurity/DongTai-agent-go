package ioReadAll

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/pkg/errors"
	"io"
	"reflect"
)

func ReadAll(r io.Reader) ([]byte, error) {

	v := reflect.ValueOf(r)
	b, e := ReadAllT(r)
	var u uintptr
	value := reflect.ValueOf(b)
	u = value.Pointer()
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(r),
		Reqs:            request.Collect(b, e),
		NeedCatch:       request.Collect(u),
		Source:          v.Type().String() == "*http.body",
		OriginClassName: "io",
		MethodName:      "ReadAll",
		ClassName:       "io",
	})
	return b, e
}

func ReadAllT(r io.Reader) ([]byte, error) {
	return []byte{}, errors.New("")
}
