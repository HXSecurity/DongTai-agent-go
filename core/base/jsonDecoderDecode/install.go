package jsonDecoderDecode

import (
	"encoding/json"
	"fmt"
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
)

func init() {
	model.HookMap["jsonDecoderDecode"] = new(JsonDecoderDecode)
}

type JsonDecoderDecode struct {
}

func (h *JsonDecoderDecode) Hook() {
	d := &json.Decoder{}
	err := gohook.HookMethod(d, "Decode", Decode, DecodeT)
	if err != nil {
		fmt.Println(err, "JsonDecoderDecode")
	} else {
		fmt.Println("JsonDecoderDecode")
	}
}

func (h *JsonDecoderDecode) UnHook() {
	d := &json.Decoder{}
	gohook.UnHookMethod(d, "JsonDecoderDecode")
}
