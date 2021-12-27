package jsonNewDecoder

import (
	"encoding/json"
	"go-agent/model"
	"go-agent/model/request"
	"go-agent/utils"
	"io"
	"reflect"
)

func NewDecoder(r io.Reader) *json.Decoder {
	v := reflect.ValueOf(r)
	d := NewDecoderT(r)
	var u uintptr
	value := reflect.ValueOf(d)
	u = value.Pointer()
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(r),
		Reqs:            utils.Collect(d),
		NeedCatch:       utils.Collect(u),
		Source:          v.Type().String() == "*http.body" && model.HookMap["ginEngineServerHTTP"] == nil,
		OriginClassName: "json",
		MethodName:      "NewDecoder",
		ClassName:       "json",
	})
	return d
}

func NewDecoderT(r io.Reader) (d *json.Decoder) {
	return
}
