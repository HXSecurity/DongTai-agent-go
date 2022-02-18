package osOpenFile

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"os"
)

func init() {
	model.HookMap["osOpenFile"] = new(OsOpenFile)
}

type OsOpenFile struct {
}

func (h *OsOpenFile) Hook() {
	err := gohook.Hook(os.OpenFile, OpenFile, OpenFileT)
	if err != nil {
		fmt.Println(err, "OsOpenFile")
	} else {
		fmt.Println("OsOpenFile")
	}
}

func (h *OsOpenFile) UnHook() {
	gohook.UnHook(os.OpenFile)
}
