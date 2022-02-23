package stringsJoin

import (
	"fmt"
)

func Join(elems []string, sep string) string {
	fmt.Println("hook到方法fmt.Sprintf")
	fmt.Println("入参", elems, sep)
	s := JoinT(elems, sep)
	fmt.Println("回参", s)
	return s
}

func JoinT(elems []string, sep string) string {
	return ""
}
