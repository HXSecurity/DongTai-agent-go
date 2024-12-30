package hook

type Mux struct {
}

func (h *Mux) GetHook() []string {
	return []string{
		"muxServeHTTP",
	}
}

func (h *Mux) HookAll() {
	Hook(h.GetHook())
}

func (h *Mux) UnHookAll() {
	UnHook(h.GetHook())
}
