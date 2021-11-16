package hookServerReq

import (
	"fmt"
	"go-agent/utils"
	"net/http"
)

// 模拟了对http包 client结构体的get的hook  此处不规范 仅为展示效果

func MyServerReq(server *http.ServeMux, w http.ResponseWriter, r *http.Request) {
	fmt.Println("请求到了\n")
	fmt.Println("入参方法", r.Method)
	fmt.Println("protocol", r.Header)
	fmt.Println("入参", r.Method)
	fmt.Println("入参", r.Method)
	fmt.Println("入参", r.Method)
	utils.CatContext()
	MyServerTempReq(server, w, r)
	return
}

func MyServerTempReq(server *http.ServeMux, w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 100; i++ {

	}
	return
}
