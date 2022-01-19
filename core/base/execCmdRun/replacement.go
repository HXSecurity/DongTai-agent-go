package execCmdRun

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"os/exec"
	"reflect"
)

func Run(cmd *exec.Cmd) error {
	e := RunT(cmd)
	//
	var u uintptr
	value := reflect.ValueOf(cmd)
	u = value.Pointer()
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(cmd.Args),
		Reqs:            utils.Collect(e),
		NeedHook:        utils.Collect(u),
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
