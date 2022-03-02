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
	"github.com/HXSecurity/DongTai-agent-go/global"
	"github.com/HXSecurity/DongTai-agent-go/hook"
)

func init() {
	g := new(hook.Gin)
	global.AllHooks = append(global.AllHooks, g)
	g.HookAll()
}
