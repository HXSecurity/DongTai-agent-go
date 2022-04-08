/*
 *  @author: Breaker
 *  @date: 2022/3/28 6:28 PM
 */

package chiServerHTTP

import (
	"fmt"

	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"github.com/go-chi/chi"
)

func init() {
	model.HookMap["chiServerHTTP"] = new(ChiServerHTTP)
}

type ChiServerHTTP struct {
}

func (h *ChiServerHTTP) Hook() {
	mux := &chi.Mux{}
	err := gohook.HookMethod(mux, "ServeHTTP", MyServer, MyServerTemp)
	if err != nil {
		fmt.Println(err, "ChiServerHTTP")
	} else {
		fmt.Println("ChiServerHTTP")
	}
}

func (h *ChiServerHTTP) UnHook() {
	mux := &chi.Mux{}
	gohook.UnHookMethod(mux, "ServeHTTP")
}
