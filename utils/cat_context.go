package utils

import (
	"bufio"
	"fmt"
	"log"
	"runtime"
	"strings"
)

//用于捕捉上下文的cat函数

func CatContext() {
	pc, file, line, _ := runtime.Caller(2)
	log.Println("调用方法等文件：" + file)
	log.Printf("调用方法所在行数:%d", line)
	f := runtime.FuncForPC(pc)
	log.Printf("数入口栈地址：%d",f.Entry())
	log.Println("调用此函数等方法：" + f.Name())
	const size = 64 << 10
	buf := make([]byte, size)
	buf = buf[:runtime.Stack(buf, true)]
	s := strings.Replace(string(buf), "\t", "", -1)
	reader := strings.NewReader(s)
	br := bufio.NewReader(reader)
	var lineArr []string
	for {
		line, _, err := br.ReadLine()
		if err != nil {
			break
		}
		lineArr = append(lineArr, string(line))
	}
	for i := len(lineArr) - 1; i > 4; i-- {
		if i%2 != 0 {
			fmt.Printf("调函数为:%s\n", lineArr[i])
		} else {
			fmt.Printf("第%d步调用文件为:%s\n", (len(lineArr)-i)/2+1, lineArr[i])
		}
	}
}
