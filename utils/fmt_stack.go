package utils

import (
	"bufio"
	"runtime"
	"strconv"
	"strings"
)

const signatureLine = 4
const callLine = 6
const callFileLine = 7

func FmtStack() (signature string, callerClass string, callerMethod string, callerLineNumber int) {
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
		if e != nil {
			break
		}
		if n == signatureLine {
			signature = l
		}
		if n == callLine {
			callArg := strings.Split(l, ".")
			callerClassArr := callArg[0 : len(callArg)-1]
			callerMethod = callArg[len(callArg)-1]
			callerClass = strings.Join(callerClassArr, ".")
		}
		if n == callFileLine {
			callFileLineBaseArr := strings.Split(l, " ")
			callFileLineArr := strings.Split(callFileLineBaseArr[0], ":")
			callerLineNumber, _ = strconv.Atoi(callFileLineArr[len(callFileLineArr)-1])
		}
		if n > callFileLine {
			break
		}
	}
	return
}
