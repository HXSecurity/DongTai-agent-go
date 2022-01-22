package urlURLQuery

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"net/http"
	"net/url"
)

func init() {
	model.HookMap["urlURLQuery"] = new(UrlURLQuery)
}

type UrlURLQuery struct {
}

func (h *UrlURLQuery) Hook() {
	var URL *url.URL
	err := gohook.HookMethod(URL, "Query", Query, QueryT)
	if err != nil {
		fmt.Println(err, "UrlURLQuery")
	} else {
		fmt.Println("UrlURLQuery")
	}
}

func (h *UrlURLQuery) UnHook() {
	var r *http.Request
	gohook.UnHookMethod(r, "UrlURLQuery")
}
