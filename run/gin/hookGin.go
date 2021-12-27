package gin

import (
	_ "go-agent/core/gin/ginContextGetPostFormArray"
	_ "go-agent/core/gin/ginContextGetPostFormMap"
	_ "go-agent/core/gin/ginContextGetQueryArray"
	_ "go-agent/core/gin/ginContextGetQueryMap"
	_ "go-agent/core/gin/ginContextParam"
	_ "go-agent/core/gin/ginContextShouldBindBodyWith"
	_ "go-agent/core/gin/ginContextShouldBindUri"
	_ "go-agent/core/gin/ginContextShouldBindWith"
	_ "go-agent/core/gin/ginEngineServerHTTP"
	_ "go-agent/core/http/httpRequestCookie"
	"go-agent/hook"
	_ "go-agent/run/base"
)

func init() {
	hook.HookFunc("ginEngineServerHTTP")
	hook.HookFunc("ginContextShouldBindWith")
	hook.HookFunc("ginContextShouldBindUri")
	hook.HookFunc("ginContextShouldBindBodyWith")
	hook.HookFunc("ginContextParam")
	hook.HookFunc("ginContextGetQueryMap")
	hook.HookFunc("ginContextGetQueryArray")
	hook.HookFunc("ginContextGetPostFormMap")
	hook.HookFunc("ginContextGetPostFormArray")
	hook.HookFunc("httpRequestCookie")
}
