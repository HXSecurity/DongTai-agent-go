package http

import (
	_ "go-agent/core/http/httpRequestCookie"
	_ "go-agent/core/http/httpRequestFormValue"
	_ "go-agent/core/http/httpServeHTTP"
	"go-agent/hook"
	_ "go-agent/run/base"
)

func init() {
	hook.HookFunc("httpServeHTTP")
}
