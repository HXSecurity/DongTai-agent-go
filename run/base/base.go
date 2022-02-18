package base

import (
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/bufioWriterWrite"
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/bufioWriterWriteString"
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/execCmdRun"
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/execCommand"
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/fmtSprintf"
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/htmlTemplateExecuteTemplate"
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/httpClientDo"
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/httpNewRequest"
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/ioReadAll"
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/jsonDecoderDecode"
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/jsonNewDecoder"
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/jsonUnmarshal"
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/osOpenFile"
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/regexpRegexpReplaceAllString"
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/runtimeConcatstrings"
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/runtimesSringtoslicebyte"
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/sqlDBQuery"
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/stringsBuilderWriteString"
	_ "github.com/HXSecurity/DongTai-agent-go/core/base/urlUrlString"

	"github.com/HXSecurity/DongTai-agent-go/service"

	"github.com/HXSecurity/DongTai-agent-go/global"
	"github.com/HXSecurity/DongTai-agent-go/hook"
)

func init() {
	global.InitViper()
	service.AgentRegister()
	hook.HookFunc("sqlDBQuery")
	hook.HookFunc("fmtSprintf")
	hook.HookFunc("jsonUnmarshal")
	hook.HookFunc("jsonDecoderDecode")
	hook.HookFunc("jsonNewDecoder")
	hook.HookFunc("runtimeConcatstrings")
	hook.HookFunc("execCommand")
	hook.HookFunc("execCmdRun")
	hook.HookFunc("bufioWriterWrite")
	hook.HookFunc("bufioWriterWriteString")
	hook.HookFunc("runtimesSringtoslicebyte")
	hook.HookFunc("htmlTemplateExecuteTemplate")
	hook.HookFunc("httpClientDo")
	hook.HookFunc("httpNewRequest")
	hook.HookFunc("stringsBuilderWriteString")
	hook.HookFunc("osOpenFile")
	hook.HookFunc("ioReadAll")
	hook.HookFunc("regexpRegexpReplaceAllString")
	hook.HookFunc("urlUrlString")
}
