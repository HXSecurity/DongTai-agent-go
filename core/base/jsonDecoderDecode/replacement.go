package jsonDecoderDecode

import (
	"encoding/json"
	"go-agent/model/request"
	"go-agent/utils"
)

func Decode(decoder *json.Decoder, v interface{}) error {
	e := DecodeT(decoder, v)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(),
		Reqs:            utils.Collect(v),
		Source:          true,
		OriginClassName: "fmt",
		MethodName:      "Sprintf",
		ClassName:       "fmt",
	})
	return e
}

func DecodeT(decoder *json.Decoder, v interface{}) error {
	return nil
}
