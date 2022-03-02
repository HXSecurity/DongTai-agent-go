package gorilla

import (
	_ "github.com/HXSecurity/DongTai-agent-go/core/gorilla/gorillaRpcServerHTTP"
	"github.com/HXSecurity/DongTai-agent-go/global"
	"github.com/HXSecurity/DongTai-agent-go/hook"
)

func init() {
	g := new(hook.Gorilla)
	global.AllHooks = append(global.AllHooks, g)
	g.HookAll()
}
