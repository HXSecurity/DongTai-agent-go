package utils

import (
	"fmt"
	"regexp"
	"runtime"
)

func CatGoroutineID() (id string) {
	buf := make([]byte, 1<<16)
	runtime.Stack(buf, false)
	str := string(buf)
	fmt.Println(str)
	r := regexp.MustCompile(`goroutine \s*(.*?)\s* \[running]:`)
	matches := r.FindAllStringSubmatch(str, -1)
	for _, v := range matches {
		id = v[1]
	}
	return id
}
