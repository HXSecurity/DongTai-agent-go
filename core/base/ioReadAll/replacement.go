package ioReadAll

import (
	"fmt"
	"github.com/pkg/errors"
	"go-agent/model/request"
	"go-agent/utils"
	"io"
	"reflect"
)

func ReadAll(r io.Reader) ([]byte, error) {
	v := reflect.ValueOf(r)
	b, e := ReadAllT(r)

	var u uintptr
	value := reflect.ValueOf(b)
	u = value.Pointer()
	if v.Type().String() == "*http.body" {
		fmt.Println(u, "in")
		fmt.Println(string(b))
	}
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(r),
		Reqs:            utils.Collect(b, e),
		NeedCatch:       utils.Collect(u),
		Source:          v.Type().String() == "*http.body",
		OriginClassName: "io",
		MethodName:      "ReadAll",
		ClassName:       "io",
	})
	return b, e
}

func ReadAllT(r io.Reader) ([]byte, error) {
	return []byte{}, errors.New("")
}
