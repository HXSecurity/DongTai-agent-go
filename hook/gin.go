package hook

type Gin struct {
}

func (g *Gin) GetHook() []string {
	return []string{
		"ginEngineServerHTTP",
		"ginContextShouldBindWith",
		"ginContextShouldBindUri",
		"ginContextShouldBindBodyWith",
		"ginContextParam",
		"ginContextGetQueryMap",
		"ginContextGetQueryArray",
		"ginContextGetPostFormMap",
		"ginContextGetPostFormArray",
		"httpRequestCookie",
	}
}

func (g *Gin) HookAll() {
	Hook(g.GetHook())
}

func (g *Gin) UnHookAll() {
	UnHook(g.GetHook())
}
