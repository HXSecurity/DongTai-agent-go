package hookSprintf

import "fmt"

func Sprintf(format string, a ...interface{}) string {
	fmt.Println("hook到方法fmt.Sprintf")
	fmt.Println("入参", format, a)
	s := SprintfT(format, a)
	fmt.Println("回参", s)
	return s
}

func SprintfT(format string, a ...interface{}) string {
	return ""
}
