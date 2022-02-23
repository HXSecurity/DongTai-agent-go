package hook

type Http struct {
}

func (h *Http) GetHook() []string {
	return []string{
		"httpServeHTTP",
	}
}

func (h *Http) HookAll() {
	Hook(h.GetHook())
}

func (h *Http) UnHookAll() {
	UnHook(h.GetHook())
}
