package httpRequestFormValue

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"net/http"
)

func init() {
	model.HookMap["httpRequestFormValue"] = new(HttpRequestFormValue)
}

type HttpRequestFormValue struct {
}

func (h *HttpRequestFormValue) Hook() {
	var r *http.Request
	err := gohook.HookMethod(r, "FormValue", FormValue, FormValueT)
	if err != nil {
		fmt.Println(err, "FormValue")
	} else {
		fmt.Println("FormValue")
	}
}

func (h *HttpRequestFormValue) UnHook() {
	var r *http.Request
	gohook.UnHookMethod(r, "FormValue")
}
