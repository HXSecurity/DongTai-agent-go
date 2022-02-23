package hook

type Gorilla struct {
}

func (g *Gorilla) GetHook() []string {
	return []string{
		"gorillaRpcServerHTTP",
	}
}

func (g *Gorilla) HookAll() {
	Hook(g.GetHook())
}

func (g *Gorilla) UnHookAll() {
	UnHook(g.GetHook())
}
