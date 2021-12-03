package utils

import (
	"fmt"
	"reflect"
)

func RangeSource(i interface{}) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
		in := v.Interface()
		RangeSource(in)
		return
	}
	if t.Kind() == reflect.Struct {
		RangeStructFiled(i)
		return
	}

	if t.Kind() == reflect.Slice {
		RangeStructSlice(i)
		return
	}

	if t.Kind() == reflect.Map {
		RangeStructMap(i)
		return
	}
}

func RangeStructFiled(i interface{}) {
	t := reflect.TypeOf(i)
	var v reflect.Value
	var value reflect.Value
	v = reflect.ValueOf(i)
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
		if value.Kind() == reflect.Struct || value.Kind() == reflect.Ptr || value.Kind() == reflect.Slice || value.Kind() == reflect.Map {
			RangeSource(value.Interface())
		}
	}
}

func RangeStructSlice(i interface{}) {
	v := reflect.ValueOf(i)
	for sl := 0; sl < v.Len(); sl++ {
		if v.Index(sl).Kind() == reflect.String {
			value := v.Index(sl)
			Type := v.Index(sl).Kind()
			pointer := GetSource(value.Interface())
			fmt.Println(value, Type, pointer)
		}
		if v.Index(sl).Kind() == reflect.Struct || v.Index(sl).Kind() == reflect.Ptr || v.Index(sl).Kind() == reflect.Slice || v.Index(sl).Kind() == reflect.Map {
			RangeSource(v.Index(sl).Interface())
		}
	}
}

func RangeStructMap(i interface{}) {
	v := reflect.ValueOf(i)
	for _, key := range v.MapKeys() {
		value := v.MapIndex(key)
		if key.Kind() == reflect.String {
			fmt.Println("map", "key", GetSource(key.Interface()))
		}

		if value.Kind() == reflect.String {
			fmt.Println("map", "value", GetSource(value.Interface()))
		}

		if key.Kind() == reflect.Struct || key.Kind() == reflect.Ptr || key.Kind() == reflect.Slice || key.Kind() == reflect.Map {
			RangeSource(key.Interface())
		}

		if value.Kind() == reflect.Struct || value.Kind() == reflect.Ptr || value.Kind() == reflect.Slice || value.Kind() == reflect.Map {
			RangeSource(value.Interface())
		}
	}
}
