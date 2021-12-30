package jsonNewDecoder

import (
	"encoding/json"
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
)

func init() {
	model.HookMap["jsonNewDecoder"] = new(JsonNewDecoder)
}

type JsonNewDecoder struct {
}

func (h *JsonNewDecoder) Hook() {
	err := gohook.Hook(json.NewDecoder, NewDecoder, NewDecoderT)
	if err != nil {
		fmt.Println(err, "JsonNewDecoder")
	} else {
		fmt.Println("JsonNewDecoder")
	}
}

func (h *JsonNewDecoder) UnHook() {
	gohook.UnHook(json.NewDecoder)
}
