package bufioWriterWriteString

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
)

func init() {
	model.HookMap["bufioWriterWriteString"] = new(BufioWriterWriteString)
}

type BufioWriterWriteString struct {
}

func (h *BufioWriterWriteString) Hook() {
	var w *bufio.Writer
	err := gohook.HookMethod(w, "WriteString", WriteString, WriteStringT)
	if err != nil {
		fmt.Println(err, "BufioWriterWriteString")
	} else {
		fmt.Println("BufioWriterWriteString")
	}
}

func (h *BufioWriterWriteString) UnHook() {
	var bt bytes.Buffer
	gohook.UnHookMethod(bt, "BufioWriterWriteString")
}
