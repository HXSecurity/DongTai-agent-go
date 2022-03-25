package httpRouter

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"github.com/go-chi/chi/v5"
)

func init() {
	model.HookMap["chiRouter"] = new(ChiRouter)
}

type ChiRouter struct {
}

func (h *ChiRouter) Hook() {
	mux := &chi.Mux{}
	err := gohook.HookMethod(mux, "ServeHTTP", ChiRouterServer, ChiRouterServerTemp)
	if err != nil {
		fmt.Println(err, "goChi")
	} else {
		fmt.Println("goChi")
	}
}

func (h *ChiRouter) UnHook() {
	mux := &chi.Mux{}
	gohook.UnHookMethod(mux, "ServeHTTP")
}
