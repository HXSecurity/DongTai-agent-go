package ioReadAll

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
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
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(r),
		Reqs:            utils.Collect(b, e),
		NeedCatch:       utils.Collect(u),
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
