package main

import (
	"encoding/json"
	"fmt"
	_ "go-agent/run/http"
	"go-agent/utils"
	"io"
	"net/http"
)

type Student struct {
	Name   string                       `json:"name"`
	Age    int                          `json:"age"`
	IdCard IdCard                       `json:"idCard"`
	Hobby  []string                     `json:"hobby"`
	MapT   map[string]map[string]string `json:"mapT"`
}

type IdCard struct {
	CardNumber string `json:"cardNumber"`
}

// 模拟用body接收一个json数据 嵌套结构体
// 将它序列化到结构体
// 其中的某些参数变化hook到具体方法

func doRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var s Student
	decoder.Decode(&s)
	s1 := fmt.Sprintf("%s", s.Name)
	fmt.Println(utils.GetSource(s.Name))
	fmt.Println(utils.GetSource(s.Name), "enter")
	m := make(map[string]string)
	m[s.Name] = s.Name
	for k, v := range m {
		fmt.Println(utils.GetSource(k), "key")
		fmt.Println(utils.GetSource(v), "value")
	}
	aa := "str2" + s.Name
	fmt.Println(s1, aa)
	w.Header().Set("test", "test")
	fmt.Fprintf(w, "service start...") //这个写入到w的是输出到客户端的 也可以用下面的 io.WriteString对象
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	//query := r.URL.Query()
	var uid string // 初始化定义变量
	if r.Method == "GET" {
		uid = r.FormValue("uid")
	} else if r.Method == "POST" {
		uid = r.PostFormValue("uid")
	}
	io.WriteString(w, "uid = "+uid)
}

func main() {
	http.HandleFunc("/test", doRequest) //   设置访问路由
	_ = http.ListenAndServe(":9090", nil)
}
