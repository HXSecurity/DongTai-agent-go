package utils

import "unsafe"

func StringAdd(s ...string) string {
	byteArr := str2bytes("")
	for _, v := range s {
		byteArr = append(byteArr, str2bytes(v)...)
	}
	return string(byteArr)
}

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
