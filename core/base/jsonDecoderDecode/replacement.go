package jsonDecoderDecode

import (
	"encoding/json"
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
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
