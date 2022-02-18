package osOpenFile

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"os"
)

func OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, e := OpenFileT(name, flag, perm)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(name, flag, perm),
		Reqs:            request.Collect(f, e),
		Source:          false,
		OriginClassName: "os",
		MethodName:      "OpenFile",
		ClassName:       "os",
	})
	return f, e
}

func OpenFileT(name string, flag int, perm os.FileMode) (*os.File, error) {
	return nil, nil
}
