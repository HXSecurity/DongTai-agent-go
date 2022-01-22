package execCommand

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"os/exec"
	"reflect"
)

func Command(name string, arg ...string) *exec.Cmd {
	e := CommandTemp(name, arg...)
	//
	var u uintptr
	value := reflect.ValueOf(e)
	u = value.Pointer()
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(name, arg),
		Reqs:            utils.Collect(e),
		NeedHook:        utils.Collect(name),
		NeedCatch:       utils.Collect(u),
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
