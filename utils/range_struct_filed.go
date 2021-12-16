package utils

import (
	"reflect"
)

func RangeSource(i interface{}, needHook *[]interface{}) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	if !v.IsValid() {
		return
	}
	if t.Kind() == reflect.Ptr {
		if v.IsNil() {
			return
		}
		t = t.Elem()
		v = v.Elem()
		in := v.Interface()
		RangeSource(in, needHook)
		return
	}
	if t.Kind() == reflect.Struct {
		RangeStructFiled(i, needHook)
		return
	}

	if t.Kind() == reflect.Slice {
		RangeStructSlice(i, needHook)
		return
	}

	if t.Kind() == reflect.Map {
		RangeStructMap(i, needHook)
		return
	}
}

func RangeStructFiled(i interface{}, needHook *[]interface{}) {
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
			if v.IsNil() {
				continue
			}
			value = v.Elem().FieldByName(k)
		}
		if value.Kind() != reflect.Struct && value.Kind() != reflect.Ptr {
			if value.Kind() == reflect.String {
				*needHook = append(*needHook, value.String())
			}
		}
		if value.Kind() == reflect.Struct || value.Kind() == reflect.Ptr || value.Kind() == reflect.Slice || value.Kind() == reflect.Map {
			if value.Kind() == reflect.Ptr && value.IsNil() {
				continue
			}
			RangeSource(value.Interface(), needHook)
		}
	}
}

func RangeStructSlice(i interface{}, needHook *[]interface{}) {
	v := reflect.ValueOf(i)
	for sl := 0; sl < v.Len(); sl++ {
		if v.Index(sl).Kind() == reflect.String {
			value := v.Index(sl)
			*needHook = append(*needHook, value.String())
		}
		if v.Index(sl).Kind() == reflect.Interface {
			value := v.Index(sl)
			switch value.Interface().(type) {
			case string:
				*needHook = append(*needHook, value.Interface())
				break
			default:
				RangeSource(value.Interface(), needHook)
				break
			}
		}

		if v.Index(sl).Kind() == reflect.Struct || v.Index(sl).Kind() == reflect.Ptr || v.Index(sl).Kind() == reflect.Slice || v.Index(sl).Kind() == reflect.Map {
			if v.Index(sl).Kind() == reflect.Ptr && v.Index(sl).IsNil() {
				continue
			}
			RangeSource(v.Index(sl).Interface(), needHook)
		}
	}
}

func RangeStructMap(i interface{}, needHook *[]interface{}) {
	v := reflect.ValueOf(i)
	for _, key := range v.MapKeys() {
		value := v.MapIndex(key)
		if key.Kind() == reflect.String {
			*needHook = append(*needHook, value.String())
		}

		if value.Kind() == reflect.String {
			*needHook = append(*needHook, value.String())
		}

		if key.Kind() == reflect.Struct || key.Kind() == reflect.Ptr || key.Kind() == reflect.Slice || key.Kind() == reflect.Map {
			if value.Kind() == reflect.Ptr && value.IsNil() {
				continue
			}
			RangeSource(key.Interface(), needHook)
		}

		if value.Kind() == reflect.Struct || value.Kind() == reflect.Ptr || value.Kind() == reflect.Slice || value.Kind() == reflect.Map {
			if value.Kind() == reflect.Ptr && value.IsNil() {
				continue
			}
			RangeSource(value.Interface(), needHook)
		}
	}
}
