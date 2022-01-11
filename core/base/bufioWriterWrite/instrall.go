package bufioWriterWrite

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
)

func init() {
	model.HookMap["bufioWriterWrite"] = new(BufioWriterWrite)
}

type BufioWriterWrite struct {
}

func (h *BufioWriterWrite) Hook() {
	var w *bufio.Writer
	err := gohook.HookMethod(w, "Write", Write, WriteT)
	if err != nil {
		fmt.Println(err, "BufioWriterWrite")
	} else {
		fmt.Println("BufioWriterWrite")
	}
}

func (h *BufioWriterWrite) UnHook() {
	var bt bytes.Buffer
	gohook.UnHookMethod(bt, "BufioWriterWrite")
}
