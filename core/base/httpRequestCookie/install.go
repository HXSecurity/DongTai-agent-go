package httpRequestCookie

import (
	"fmt"
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
	"net/http"
)

func init() {
	model.HookMap["httpRequestCookie"] = new(HttpRequestCookie)
}

type HttpRequestCookie struct {
}

func (h *HttpRequestCookie) Hook() {
	var r *http.Request
	err := gohook.HookMethod(r, "Cookie", Cookie, CookieT)
	if err != nil {
		fmt.Println(err, "Cookie")
	} else {
		fmt.Println("Cookie")
	}
}

func (h *HttpRequestCookie) UnHook() {
	var r *http.Request
	gohook.UnHookMethod(r, "Cookie")
}
