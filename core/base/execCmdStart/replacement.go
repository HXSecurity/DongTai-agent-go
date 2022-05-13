package execCmdStart

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"os/exec"
	"reflect"
)

func Start(cmd *exec.Cmd) error {
	e := StartT(cmd)
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
		MethodName:      "Start",
		ClassName:       "exec.(*Cmd)",
	})
	return e
}

func StartT(cmd *exec.Cmd) error {
	return nil
}
