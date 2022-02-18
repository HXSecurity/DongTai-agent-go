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
	var WriteMap = make(map[string]bool)
	WriteMap["github.com/HXSecurity/DongTai-agent-go/core/httpRouter/httpRouter.MyHttpRouterServer.func1"] = true
	WriteMap["os.openDir"] = true
	WriteMap["io/fs.(*PathError).Error"] = true
	WriteMap["syscall.UTF16FromString"] = true
	WriteMap["os/exec.findExecutable"] = true
	if !WriteMap[utils.LoadFunc(3)] {
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
		request.FmtHookPool(request.PoolReq{
			Args:            request.Collect(a),
			ArgsStr:         ArgStr,
			Reqs:            request.Collect(e),
			NeedHook:        NeedHook,
			NeedCatch:       request.Collect(e),
			Source:          false,
			OriginClassName: "runtime",
			MethodName:      "concatstrings",
			ClassName:       "runtime",
		})
	}
	return e
}

func concatstringsT(buf *tmpBuf, a []string) string {
	return ""
}
