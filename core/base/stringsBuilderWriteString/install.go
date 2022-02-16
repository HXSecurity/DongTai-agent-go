package stringsBuilderWriteString

import (
	"bytes"
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"strings"
)

func init() {
	model.HookMap["stringsBuilderWriteString"] = new(BufferWriteString)
}

type BufferWriteString struct {
}

func (h *BufferWriteString) Hook() {
	var bt *strings.Builder
	err := gohook.HookMethod(bt, "WriteString", WriteString, WriteStringT)
	if err != nil {
		fmt.Println(err, "stringsBuilderWriteString")
	} else {
		fmt.Println("stringsBuilderWriteString")
	}
}

func (h *BufferWriteString) UnHook() {
	var bt bytes.Buffer
	gohook.UnHookMethod(bt, "stringsBuilderWriteString")
}
