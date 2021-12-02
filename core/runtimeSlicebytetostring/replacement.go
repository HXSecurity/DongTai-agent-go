package runtimeSlicebytetostring

import (
	"fmt"
	"go-agent/utils"
	_ "unsafe"
)

const tmpStringBufSize = 32

type tmpBuf [tmpStringBufSize]byte

//go:linkname slicebytetostring runtime.slicebytetostring
func slicebytetostring(buf *tmpBuf, ptr *byte, n int) (str string)

func slicebytetostringR(buf *tmpBuf, ptr *byte, n int) (str string) {

	//fmt.Println("Hook到了slicebytetostring方法")
	//fmt.Println("入参为", buf,ptr,n)
	e := slicebytetostringT(buf, ptr, n)
	if utils.IsHook("go-agent/core/httpServeHTTP.MyServer", 7) {
		fmt.Println("在http路由入口返回值为", e)
	}
	//fmt.Println("返回值", e)
	return e
}

func slicebytetostringT(buf *tmpBuf, ptr *byte, n int) (str string) {
	return ""
}
