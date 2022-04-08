/*
 *  @author: Breaker
 *  @date: 2022/3/29 10:23 AM
 */

package chiURLParam

import (
	"fmt"

	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"github.com/go-chi/chi"
)

func init() {
	model.HookMap["chiURLParam"] = new(ChiURLParam)
}

type ChiURLParam struct {
}

func (h *ChiURLParam) Hook() {
	err := gohook.Hook(chi.URLParam, chiURLParam, chiURLParamTemp)
	if err != nil {
		fmt.Println(err, "ChiURLParam")
	} else {
		fmt.Println("ChiURLParam")
	}
}

func (h *ChiURLParam) UnHook() {
	gohook.UnHook(chi.URLParam)
}
