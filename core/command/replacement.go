package hookExecCommand

import (
	"fmt"
	"os/exec"
)

func Command(name string, arg ...string) *exec.Cmd {
	fmt.Printf("Hook到了%s\n方法", "Command")
	fmt.Printf("入参为%s,%s\n", name, arg)
	e := CommandTemp(name, arg...)
	fmt.Printf("返回值%s\n", e)
	return e
}

func CommandTemp(name string, arg ...string) *exec.Cmd {
	return &exec.Cmd{}
}
