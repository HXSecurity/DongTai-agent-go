package hook

type Http struct {
}

func (h *Http) GetHook() []string {
	return []string{
		"httpServeHTTP",
		"httpRequestFormValue",
		"httpServeHTTP",
		"urlURLQuery",
	}
}

func (h *Http) HookAll() {
	Hook(h.GetHook())
}

func (h *Http) UnHookAll() {
	UnHook(h.GetHook())
}
