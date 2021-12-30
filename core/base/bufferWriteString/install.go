package bufferWriteString

import (
	"bytes"
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
)

func init() {
	model.HookMap["bufferWriteString"] = new(BufferWriteString)
}

type BufferWriteString struct {
}

func (h *BufferWriteString) Hook() {
	var bt bytes.Buffer
	err := gohook.HookMethod(bt, "WriteString", WriteString, WriteStringT)
	if err != nil {
		fmt.Println(err, "BufferWriteString")
	} else {
		fmt.Println("BufferWriteString")
	}
}

func (h *BufferWriteString) UnHook() {
	var bt bytes.Buffer
	gohook.UnHookMethod(bt, "WriteString")
}
