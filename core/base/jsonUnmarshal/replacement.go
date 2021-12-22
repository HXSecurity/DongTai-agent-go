package jsonUnmarshal

import (
	"go-agent/model/request"
	"go-agent/utils"
)

func Unmarshal(data []byte, v interface{}) error {
	e := UnmarshalT(data, v)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(data),
		Reqs:            utils.Collect(v),
		Source:          true,
		OriginClassName: "fmt",
		MethodName:      "Sprintf",
		ClassName:       "fmt",
	})
	return e
}

func UnmarshalT(data []byte, v interface{}) error {
	return nil
}
