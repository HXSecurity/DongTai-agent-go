package urlUrlString

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"net/url"
)

func init() {
	model.HookMap["urlUrlString"] = new(UrlUrlString)
}

type UrlUrlString struct {
}

func (h *UrlUrlString) Hook() {
	var url *url.URL
	err := gohook.HookMethod(url, "String", String, StringT)
	if err != nil {
		fmt.Println(err, "UrlUrlString")
	} else {
		fmt.Println("UrlUrlString")
	}
}

func (h *UrlUrlString) UnHook() {
	var url *url.URL
	gohook.UnHookMethod(url, "UrlUrlString")
}
