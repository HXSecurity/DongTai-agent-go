package runtimeConcatstrings

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	_ "unsafe"
)

const tmpStringBufSize = 32

type tmpBuf [tmpStringBufSize]byte

//go:linkname concatstrings runtime.concatstrings
func concatstrings(buf *tmpBuf, a []string) string

func concatstringsR(buf *tmpBuf, a []string) string {
	e := concatstringsT(buf, a)
	var ArgStr = "["
	for k, v := range a {
		ArgStr = utils.StringAdd(ArgStr, v)
		if k != len(a)-1 {
			ArgStr = utils.StringAdd(ArgStr, ",")
		} else {
			ArgStr = utils.StringAdd(ArgStr, "]")
		}
	}
	NeedHook := make([]interface{}, len(a))
	for i, v := range a {
		NeedHook[i] = v
	}
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(a),
		ArgsStr:         ArgStr,
		Reqs:            utils.Collect(e),
		NeedHook:        NeedHook,
		NeedCatch:       utils.Collect(e),
		Source:          false,
		OriginClassName: "runtime",
		MethodName:      "concatstrings",
		ClassName:       "runtime",
	})

	return e
}

func concatstringsT(buf *tmpBuf, a []string) string {
	return ""
}
