package builderWriteString

import (
	"bytes"
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
	"strings"
)

func init() {
	model.HookMap["builderWriteString"] = new(BuilderWriteString)
}

type BuilderWriteString struct {
}

func (h *BuilderWriteString) Hook() {
	var bt strings.Builder
	gohook.HookMethod(bt, "WriteString", WriteString, WriteStringT)
}

func (h *BuilderWriteString) UnHook() {
	var bt bytes.Buffer
	gohook.UnHookMethod(bt, "WriteString")
}
