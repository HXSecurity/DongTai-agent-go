package gorilla

import (
	_ "github.com/HXSecurity/DongTai-agent-go/core/gorilla/gorillaRpcServerHTTP"
	"github.com/HXSecurity/DongTai-agent-go/hook"
)

func init() {
	hook.HookFunc("gorillaRpcServerHTTP")
}
