package jsonUnmarshal

import (
	"go-agent/model/request"
	"go-agent/utils"
	"reflect"
)

func Unmarshal(data []byte, v interface{}) error {
	var u uintptr
	value := reflect.ValueOf(data)
	u = value.Pointer()
	e := UnmarshalT(data, v)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(data),
		Reqs:            utils.Collect(v),
		NeedHook:        utils.Collect(u),
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
