package httpNewRequest

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"net/http"
)

func init() {
	model.HookMap["httpNewRequest"] = new(HttpNewRequest)
}

type HttpNewRequest struct {
}

func (h *HttpNewRequest) Hook() {
	err := gohook.Hook(http.NewRequest, NewRequest, NewRequestR)
	if err != nil {
		fmt.Println(err, "HttpNewRequest")
	} else {
		fmt.Println("HttpNewRequest")
	}
}

func (h *HttpNewRequest) UnHook() {
	gohook.UnHook(http.NewRequest)
}
