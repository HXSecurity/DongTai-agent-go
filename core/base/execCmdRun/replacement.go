package execCmdRun

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"os/exec"
	"reflect"
)

func Run(cmd *exec.Cmd) error {
	e := RunT(cmd)
	//
	var u uintptr
	value := reflect.ValueOf(cmd)
	u = value.Pointer()
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(cmd.Args),
		Reqs:            request.Collect(e),
		NeedHook:        request.Collect(u),
		Source:          false,
		OriginClassName: "exec.(*Cmd)",
		MethodName:      "Run",
		ClassName:       "exec.(*Cmd)",
	})
	return e
}

func RunT(cmd *exec.Cmd) error {
	return nil
}
