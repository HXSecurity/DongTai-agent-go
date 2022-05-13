package execCommand

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"os/exec"
	"reflect"
)

func Command(name string, arg ...string) *exec.Cmd {
	e := CommandTemp(name, arg...)
	//
	var u uintptr
	value := reflect.ValueOf(e)
	u = value.Pointer()
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(name, arg),
		Reqs:            request.Collect(e),
		NeedCatch:       request.Collect(u),
		Source:          false,
		OriginClassName: "exec",
		MethodName:      "Command",
		ClassName:       "exec",
	})
	return e
}

func CommandTemp(name string, arg ...string) *exec.Cmd {
	return &exec.Cmd{}
}
