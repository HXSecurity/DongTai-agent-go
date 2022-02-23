package http

import (
	_ "github.com/HXSecurity/DongTai-agent-go/core/http/httpRequestCookie"
	_ "github.com/HXSecurity/DongTai-agent-go/core/http/httpRequestFormValue"
	_ "github.com/HXSecurity/DongTai-agent-go/core/http/httpServeHTTP"
	"github.com/HXSecurity/DongTai-agent-go/hook"
	_ "github.com/HXSecurity/DongTai-agent-go/run/base"
)

func init() {
	h := new(hook.Http)
	h.HookAll()
}
