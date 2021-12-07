package runtimeConcatstrings

import (
	"fmt"
	"go-agent/utils"
	_ "unsafe"
)

const tmpStringBufSize = 32

type tmpBuf [tmpStringBufSize]byte

//go:linkname concatstrings runtime.concatstrings
func concatstrings(buf *tmpBuf, a []string) string

func concatstringsR(buf *tmpBuf, a []string) string {
	fmt.Println("Hook到了concatstrings方法")
	fmt.Println("入参为", buf, a)
	for _, v := range a {
		fmt.Println("source:", utils.GetSource(v), "原始value", v)
	}
	e := concatstringsT(buf, a)
	fmt.Println("target:", utils.GetSource(e), "合并后", e)
	fmt.Println("返回值", e)
	return e
}

func concatstringsT(buf *tmpBuf, a []string) string {
	return ""
}
