package httpRequestFormValue

import (
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
	"net/http"
)

func init() {
	model.HookMap["httpRequestFormValue"] = new(HttpRequestFormValue)
}

type HttpRequestFormValue struct {
}

func (h *HttpRequestFormValue) Hook() {
	var r *http.Request
	gohook.HookMethod(r, "FormValue", FormValue, FormValueT)
}

func (h *HttpRequestFormValue) UnHook() {
	var r *http.Request
	gohook.UnHookMethod(r, "FormValue")
}
