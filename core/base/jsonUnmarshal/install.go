package jsonUnmarshal

import (
	"encoding/json"
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
)

func init() {
	model.HookMap["jsonUnmarshal"] = new(JsonUnmarshal)
}

type JsonUnmarshal struct {
}

func (h *JsonUnmarshal) Hook() {
	err := gohook.Hook(json.Unmarshal, Unmarshal, UnmarshalT)
	if err != nil {
		fmt.Println(err, "JsonUnmarshal")
	} else {
		fmt.Println("JsonUnmarshal")
	}
}

func (h *JsonUnmarshal) UnHook() {
	gohook.UnHook(json.Unmarshal)
}
