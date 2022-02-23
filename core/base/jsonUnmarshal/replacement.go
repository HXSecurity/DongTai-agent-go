package jsonUnmarshal

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"reflect"
)

func Unmarshal(data []byte, v interface{}) error {

	var u uintptr
	value := reflect.ValueOf(data)
	u = value.Pointer()
	e := UnmarshalT(data, v)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(data),
		Reqs:            request.Collect(v),
		NeedHook:        request.Collect(u),
		Source:          false,
		OriginClassName: "json",
		MethodName:      "Unmarshal",
		ClassName:       "json",
	})
	return e
}

func UnmarshalT(data []byte, v interface{}) error {
	return nil
}
