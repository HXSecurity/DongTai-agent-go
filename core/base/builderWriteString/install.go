package builderWriteString

import (
	"bytes"
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"strings"
)

func init() {
	model.HookMap["builderWriteString"] = new(BuilderWriteString)
}

type BuilderWriteString struct {
}

func (h *BuilderWriteString) Hook() {
	var bt strings.Builder
	err := gohook.HookMethod(bt, "WriteString", WriteString, WriteStringT)
	if err != nil {
		fmt.Println(err, "BuilderWriteString")
	} else {
		fmt.Println("BuilderWriteString")
	}
}

func (h *BuilderWriteString) UnHook() {
	var bt bytes.Buffer
	gohook.UnHookMethod(bt, "WriteString")
}
