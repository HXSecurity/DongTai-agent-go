package gin

import (
	_ "github.com/HXSecurity/DongTai-agent-go/core/gorm/gormDBExec"
	_ "github.com/HXSecurity/DongTai-agent-go/core/gorm/gormDBGroup"
	_ "github.com/HXSecurity/DongTai-agent-go/core/gorm/gormDBHaving"
	_ "github.com/HXSecurity/DongTai-agent-go/core/gorm/gormDBOrder"
	_ "github.com/HXSecurity/DongTai-agent-go/core/gorm/gormDBPluck"
	_ "github.com/HXSecurity/DongTai-agent-go/core/gorm/gormDBRaw"
	_ "github.com/HXSecurity/DongTai-agent-go/core/gorm/gormDBSelect"
	"github.com/HXSecurity/DongTai-agent-go/hook"
)

func init() {
	hook.HookFunc("gormDBOrder")
	hook.HookFunc("gormDBExec")
	hook.HookFunc("gormDBGroup")
	hook.HookFunc("gormDBHaving")
	hook.HookFunc("gormDBPluck")
	hook.HookFunc("gormDBRaw")
	hook.HookFunc("gormDBSelect")
}
