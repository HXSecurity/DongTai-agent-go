package bufferWriteString

import (
	"bytes"
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
)

func init() {
	model.HookMap["bufferWriteString"] = new(BufferWriteString)
}

type BufferWriteString struct {
}

func (h *BufferWriteString) Hook() {
	var bt bytes.Buffer
	gohook.HookMethod(bt, "WriteString", WriteString, WriteStringT)
}

func (h *BufferWriteString) UnHook() {
	var bt bytes.Buffer
	gohook.UnHookMethod(bt, "WriteString")
}
