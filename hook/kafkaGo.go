package hook

type KafkaGo struct {
}

func (g *KafkaGo) GetHook() []string {
	return []string{
		"kafkaGoWriter",
	}
}

func (g *KafkaGo) HookAll() {
	Hook(g.GetHook())
}

func (g *KafkaGo) UnHookAll() {
	UnHook(g.GetHook())
}
