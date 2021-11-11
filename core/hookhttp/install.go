package hookhttp

import (
	"fmt"
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
	"net/http"
)

func init(){
	model.HookMap["http"] = new(HookHttp)
}



type HookHttp struct {

}

func (h *HookHttp)Hook(){
	c := &http.Client{}
	err:= gohook.HookMethod(c,"Get", MyGet, MyGetTemp)
	fmt.Println(err)
}

func (h *HookHttp)UnHook(){
	c := &http.Client{}
	gohook.UnHookMethod(c,"Get")
}


