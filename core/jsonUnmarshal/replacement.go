package jsonUnmarshal

import (
	"fmt"
	"reflect"
)

func Unmarshal(data []byte, v interface{}) error {
	fmt.Println(reflect.ValueOf(data).Pointer(), "Unmarshal")
	e := UnmarshalT(data, v)
	fmt.Println(reflect.ValueOf(v).Pointer(), "UnmarshalR")
	return e
}

func UnmarshalT(data []byte, v interface{}) error {
	return nil
}
