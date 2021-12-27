package jsonDecoderDecode

import (
	"encoding/json"
	"go-agent/model/request"
	"go-agent/utils"
	"reflect"
)

func Decode(decoder *json.Decoder, v interface{}) error {
	var u uintptr
	value := reflect.ValueOf(decoder)
	u = value.Pointer()
	e := DecodeT(decoder, v)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(decoder),
		Reqs:            utils.Collect(v),
		NeedHook:        utils.Collect(u),
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
