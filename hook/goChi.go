package hook

type ChiRouter struct {
}

func (h *ChiRouter) GetHook() []string {
	return []string{
		"chiRouter",
		"httpRequestFormValue",
		"httpRequestCookie",
		"urlURLQuery",
	}
}

func (h *ChiRouter) HookAll() {
	Hook(h.GetHook())
}

func (h *ChiRouter) UnHookAll() {
	UnHook(h.GetHook())
}
