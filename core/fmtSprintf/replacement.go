package fmtSprintf

import (
	"fmt"
	"go-agent/utils"
)

func Sprintf(format string, a ...interface{}) string {
	for _, v := range a {
		switch v.(type) {
		case string:
			fmt.Println(utils.GetSource(v))
		}
	}
	s := SprintfT(format, a)
	return s
}

func SprintfT(format string, a ...interface{}) string {
	return ""
}
