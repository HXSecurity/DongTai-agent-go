package runtimesSringtoslicebyte

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"reflect"
	_ "unsafe"
)

const tmpStringBufSize = 32

type tmpBuf [tmpStringBufSize]byte

//go:linkname stringtoslicebyte runtime.stringtoslicebyte
func stringtoslicebyte(buf *tmpBuf, s string) []byte

func stringtoslicebyteR(buf *tmpBuf, s string) []byte {
	b := stringtoslicebyteT(buf, s)

	var u uintptr
	value := reflect.ValueOf(b)
	u = value.Pointer()

	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(s),
		Reqs:            utils.Collect(b),
		NeedCatch:       utils.Collect(u),
		Source:          false,
		OriginClassName: "runtime",
		MethodName:      "stringtoslicebyte",
		ClassName:       "runtime",
	})

	return b
}

func stringtoslicebyteT(buf *tmpBuf, s string) []byte {
	return []byte{}
}
