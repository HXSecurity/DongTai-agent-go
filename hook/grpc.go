package hook

type Grpc struct {
}

func (g *Grpc) GetHook() []string {
	return []string{
		"grpcClientConn",
		"grpcNewServer",
		"grpcUnaryInterceptor",
	}
}

func (g *Grpc) HookAll() {
	Hook(g.GetHook())
}

func (g *Grpc) UnHookAll() {
	UnHook(g.GetHook())
}
