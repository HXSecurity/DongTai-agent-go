package httpRouter

import (
	_ "go-agent/core/httpRouter/httpRouter"
	"go-agent/hook"
	_ "go-agent/run/base"
)

func init() {
	hook.HookFunc("httpRouter")
}
