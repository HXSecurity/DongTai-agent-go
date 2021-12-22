package gin

import (
	_ "go-agent/core/gorm/gormDBOrder"
	"go-agent/hook"
)

func init() {
	hook.HookFunc("gormDBOrder")
}
