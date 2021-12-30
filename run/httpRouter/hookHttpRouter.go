package httpRouter

import (
	_ "github.com/HXSecurity/DongTai-agent-go/core/http/httpRequestCookie"
	_ "github.com/HXSecurity/DongTai-agent-go/core/http/httpRequestFormValue"
	_ "github.com/HXSecurity/DongTai-agent-go/core/httpRouter/httpRouter"
	"github.com/HXSecurity/DongTai-agent-go/hook"
	_ "github.com/HXSecurity/DongTai-agent-go/run/base"
)

func init() {
	hook.HookFunc("httpRouter")
	hook.HookFunc("httpRequestFormValue")
	hook.HookFunc("httpRequestCookie")
}
