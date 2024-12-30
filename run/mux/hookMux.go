package http

import (
	_ "github.com/HXSecurity/DongTai-agent-go/core/mux/muxServerHTTP"
	"github.com/HXSecurity/DongTai-agent-go/global"
	"github.com/HXSecurity/DongTai-agent-go/hook"
)

func init() {
	h := new(hook.Mux)
	global.AllHooks = append(global.AllHooks, h)
	h.HookAll()
}
