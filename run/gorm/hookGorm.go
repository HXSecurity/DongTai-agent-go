package gin

import (
	_ "go-agent/core/gorm/gormDBExec"
	_ "go-agent/core/gorm/gormDBGroup"
	_ "go-agent/core/gorm/gormDBHaving"
	_ "go-agent/core/gorm/gormDBOrder"
	_ "go-agent/core/gorm/gormDBPluck"
	_ "go-agent/core/gorm/gormDBRaw"
	_ "go-agent/core/gorm/gormDBSelect"
	"go-agent/hook"
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
