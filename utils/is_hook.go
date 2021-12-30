package utils

import (
	"bufio"
	"runtime"
	"strings"
	"unsafe"
)

func IsHook(string2 string, line int) (flag bool) {
	l := 1 << 10
	buf := make([]byte, l)
	runtime.Stack(buf, false)
	s := byteToString(buf)
	r := strings.NewReader(s)
	reader := bufio.NewReader(r)
	n := 0
	for {
		l, e := reader.ReadString('\n')
		n++
		if n == line {
			if strings.Index(l, string2) > -1 {
				return true
			}
		}
		if e != nil {
			return flag
		}
	}
}

func byteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
