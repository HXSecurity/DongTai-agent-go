package jsonUnmarshal

import (
	"fmt"
	"go-agent/utils"
	"reflect"
)

func Unmarshal(data []byte, v interface{}) error {
	fmt.Println(reflect.ValueOf(data).Pointer(), "Unmarshal")
	e := UnmarshalT(data, v)
	fmt.Println(reflect.ValueOf(v).Pointer(), "UnmarshalR")
	utils.RangeStructFiled(v)
	return e
}

func UnmarshalT(data []byte, v interface{}) error {
	return nil
}
