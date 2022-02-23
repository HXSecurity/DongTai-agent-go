package jsonDecoderDecode

import (
	"encoding/json"
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"reflect"
)

func Decode(decoder *json.Decoder, v interface{}) error {
	var u uintptr
	value := reflect.ValueOf(decoder)
	u = value.Pointer()
	e := DecodeT(decoder, v)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(decoder),
		Reqs:            request.Collect(v),
		NeedHook:        request.Collect(u),
		Source:          false,
		OriginClassName: "json.(*Decoder)",
		MethodName:      "Decode",
		ClassName:       "json.(*Decoder)",
	})
	return e
}

func DecodeT(decoder *json.Decoder, v interface{}) error {
	return nil
}
