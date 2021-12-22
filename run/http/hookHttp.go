package http

import (
	_ "go-agent/core/base/httpServeHTTP"
	"go-agent/hook"
	_ "go-agent/run/base"
)

func init() {
	hook.HookFunc("httpServeHTTP")
}
