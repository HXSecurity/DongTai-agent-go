/*
 *  @author: YangRui
 *  @date: 2022/4/8 11:21 AM
 */

package chi

import (
	"github.com/HXSecurity/DongTai-agent-go/global"
	"github.com/HXSecurity/DongTai-agent-go/hook"

	_ "github.com/HXSecurity/DongTai-agent-go/core/chi/chiServerHTTP"
	_ "github.com/HXSecurity/DongTai-agent-go/core/chi/chiURLParam"
	_ "github.com/HXSecurity/DongTai-agent-go/core/http/httpRequestCookie"
)

func init() {
	g := new(hook.Chi)
	global.AllHooks = append(global.AllHooks, g)
	g.HookAll()
}
