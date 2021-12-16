package httpRequestCookie

import (
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
	gohook.HookMethod(r, "Cookie", Cookie, CookieT)
}

func (h *HttpRequestCookie) UnHook() {
	var r *http.Request
	gohook.UnHookMethod(r, "Cookie")
}
