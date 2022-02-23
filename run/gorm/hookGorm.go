package gorm

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
	g := new(hook.Gorm)
	g.HookAll()
}
