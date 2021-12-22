package gin

import (
	_ "go-agent/core/gin/ginEngineServerHTTP"
	"go-agent/hook"
	_ "go-agent/run/base"
)

func init() {
	hook.HookFunc("ginEngineServerHTTP")
}
