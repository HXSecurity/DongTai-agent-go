package httpRouter

import (
	_ "github.com/HXSecurity/DongTai-agent-go/core/http/httpRequestCookie"
	_ "github.com/HXSecurity/DongTai-agent-go/core/http/httpRequestFormValue"
	_ "github.com/HXSecurity/DongTai-agent-go/core/http/urlURLQuery"
	_ "github.com/HXSecurity/DongTai-agent-go/core/httpRouter/httpRouter"
	"github.com/HXSecurity/DongTai-agent-go/global"
	"github.com/HXSecurity/DongTai-agent-go/hook"
)

func init() {
	h := new(hook.ChiRouter)
	global.AllHooks = append(global.AllHooks, h)
	h.HookAll()
}
