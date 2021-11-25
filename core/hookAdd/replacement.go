package hookAdd

import (
	"fmt"
	_ "unsafe"
)

const tmpStringBufSize = 32

type tmpBuf [tmpStringBufSize]byte

//go:linkname concatstrings runtime.concatstrings
func concatstrings(buf *tmpBuf, a []string) string

func concatstringsR(buf *tmpBuf, a []string) string {
	fmt.Println("Hook到了concatstrings方法")
	fmt.Println("入参为", buf, a)
	e := concatstringsT(buf, a)
	fmt.Println("返回值", e)
	return e
}

func concatstringsT(buf *tmpBuf, a []string) string {
	return ""
}
