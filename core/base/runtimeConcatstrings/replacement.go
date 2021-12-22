package runtimeConcatstrings

import (
	"go-agent/model/request"
	"go-agent/utils"
	_ "unsafe"
)

const tmpStringBufSize = 32

type tmpBuf [tmpStringBufSize]byte

//go:linkname concatstrings runtime.concatstrings
func concatstrings(buf *tmpBuf, a []string) string

func concatstringsR(buf *tmpBuf, a []string) string {
	e := concatstringsT(buf, a)
	if utils.IsHook("reflect.(*rtype).ptrTo", 6) {
		return e
	}
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(a),
		Reqs:            utils.Collect(e),
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
