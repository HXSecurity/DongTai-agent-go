package jsonUnmarshal

import (
	"encoding/json"
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
)

func init() {
	model.HookMap["jsonUnmarshal"] = new(JsonUnmarshal)
}

type JsonUnmarshal struct {
}

func (h *JsonUnmarshal) Hook() {
	gohook.Hook(json.Unmarshal, Unmarshal, UnmarshalT)
}

func (h *JsonUnmarshal) UnHook() {
	gohook.UnHook(json.Unmarshal)
}
