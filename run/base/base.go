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
	"github.com/HXSecurity/DongTai-agent-go/hook"
	"github.com/HXSecurity/DongTai-agent-go/service"
	"github.com/google/uuid"

	"github.com/HXSecurity/DongTai-agent-go/global"
)

func init() {
	b := new(hook.Base)
	global.AllHooks = append(global.AllHooks, b)
	global.InitViper()
	_ = service.AgentRegister()
	b.HookAll()
	global.TraceId = uuid.New().String() + "-" + uuid.New().String()
}
