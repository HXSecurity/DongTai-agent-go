package hook

type Http struct {
}

func (h *Http) GetHook() []string {
	return []string{
		"httpHeaderGet",
		"httpRequestCookie",
		"httpServeHTTP",
		"httpRequestFormValue",
		"urlURLQuery",
	}
}

func (h *Http) HookAll() {
	Hook(h.GetHook())
}

func (h *Http) UnHookAll() {
	UnHook(h.GetHook())
}
