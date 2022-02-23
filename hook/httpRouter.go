package hook

type HttpRouter struct {
}

func (h *HttpRouter) GetHook() []string {
	return []string{
		"httpRouter",
		"httpRequestFormValue",
		"httpRequestCookie",
		"urlURLQuery",
	}
}

func (h *HttpRouter) HookAll() {
	Hook(h.GetHook())
}

func (h *HttpRouter) UnHookAll() {
	UnHook(h.GetHook())
}
