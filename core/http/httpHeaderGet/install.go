package httpHeaderGet

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"net/http"
)

func init() {
	model.HookMap["httpHeaderGet"] = new(HttpHeaderGet)
}

type HttpHeaderGet struct {
}

func (h *HttpHeaderGet) Hook() {
	var header http.Header
	err := gohook.HookMethod(header, "Get", Get, GetT)
	if err != nil {
		fmt.Println(err, "HttpHeaderGet")
	} else {
		fmt.Println("HttpHeaderGet")
	}
}

func (h *HttpHeaderGet) UnHook() {
	var header *http.Header
	gohook.UnHookMethod(header, "Get")
}
