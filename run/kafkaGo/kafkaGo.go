package kafkaGo

import (
	_ "github.com/HXSecurity/DongTai-agent-go/core/kafkaGo/kafkaWriter"
	"github.com/HXSecurity/DongTai-agent-go/global"
	"github.com/HXSecurity/DongTai-agent-go/hook"
)

func init() {
	g := new(hook.KafkaGo)
	global.AllHooks = append(global.AllHooks, g)
	g.HookAll()
}
