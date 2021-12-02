package utils

import (
	"fmt"
	"reflect"
)

func RangeStructFiled(i interface{}) {
	t := reflect.TypeOf(i)
	var v reflect.Value
	var value reflect.Value
	v = reflect.ValueOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		fmt.Println("非结构体不可遍历")
		return
	}
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		k := t.Field(i).Name
		if v.Kind() == reflect.Struct {
			value = v.FieldByName(k)
		}
		if v.Kind() == reflect.Ptr {
			value = v.Elem().FieldByName(k)
		}
		ty := value.Type()
		if value.Kind() != reflect.Struct && value.Kind() != reflect.Ptr {
			fmt.Println(k, value, ty, value.Kind())
			if value.Kind() == reflect.String {
				fmt.Println(GetSource(value.Interface()))
			}
		}
		if value.Kind() == reflect.Struct || value.Kind() == reflect.Ptr {
			RangeStructFiled(value.Interface())
		}
	}
}
