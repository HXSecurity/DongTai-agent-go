package grpc

import (
	_ "github.com/HXSecurity/DongTai-agent-go/core/grpc/clientConn"
	_ "github.com/HXSecurity/DongTai-agent-go/core/grpc/newServer"
	_ "github.com/HXSecurity/DongTai-agent-go/core/grpc/unaryInterceptor"
	"github.com/HXSecurity/DongTai-agent-go/global"
	"github.com/HXSecurity/DongTai-agent-go/hook"
)

func init() {
	g := new(hook.Grpc)
	global.AllHooks = append(global.AllHooks, g)
	g.HookAll()
}
