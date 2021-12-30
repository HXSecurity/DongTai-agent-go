package gin

import (
	_ "github.com/HXSecurity/DongTai-agent-go/core/gin/ginContextGetPostFormArray"
	_ "github.com/HXSecurity/DongTai-agent-go/core/gin/ginContextGetPostFormMap"
	_ "github.com/HXSecurity/DongTai-agent-go/core/gin/ginContextGetQueryArray"
	_ "github.com/HXSecurity/DongTai-agent-go/core/gin/ginContextGetQueryMap"
	_ "github.com/HXSecurity/DongTai-agent-go/core/gin/ginContextParam"
	_ "github.com/HXSecurity/DongTai-agent-go/core/gin/ginContextShouldBindBodyWith"
	_ "github.com/HXSecurity/DongTai-agent-go/core/gin/ginContextShouldBindUri"
	_ "github.com/HXSecurity/DongTai-agent-go/core/gin/ginContextShouldBindWith"
	_ "github.com/HXSecurity/DongTai-agent-go/core/gin/ginEngineServerHTTP"
	_ "github.com/HXSecurity/DongTai-agent-go/core/http/httpRequestCookie"
	"github.com/HXSecurity/DongTai-agent-go/hook"
	_ "github.com/HXSecurity/DongTai-agent-go/run/base"
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
