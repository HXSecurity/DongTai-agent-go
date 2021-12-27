package httpRouter

import (
	_ "go-agent/core/http/httpRequestCookie"
	_ "go-agent/core/http/httpRequestFormValue"
	_ "go-agent/core/httpRouter/httpRouter"
	"go-agent/hook"
	_ "go-agent/run/base"
)

func init() {
	hook.HookFunc("httpRouter")
	hook.HookFunc("httpRequestFormValue")
	hook.HookFunc("httpRequestCookie")
}
