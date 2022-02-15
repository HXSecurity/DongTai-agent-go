package httpClientDo

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"net/http"
)

func init() {
	model.HookMap["httpClientDo"] = new(HttpClientDo)
}

type HttpClientDo struct {
}

func (h *HttpClientDo) Hook() {
	client := &http.Client{}
	err := gohook.HookMethod(client, "Do", Do, DoR)
	if err != nil {
		fmt.Println(err, "HttpClientDo")
	} else {
		fmt.Println("HttpClientDo")
	}
}

func (h *HttpClientDo) UnHook() {
	client := &http.Client{}
	gohook.UnHookMethod(client, "Do")
}
