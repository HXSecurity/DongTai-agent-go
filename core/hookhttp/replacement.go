package hookhttp

import (
	"fmt"
	"go-agent/utils"
	"net/http"
)

// 模拟了对http包 client结构体的get的hook  此处不规范 仅为展示效果

func MyGet(client *http.Client,url string) (resp *http.Response, err error) {
	fmt.Println("hook到了 *http.Client 的 Get 方法")
	fmt.Println("入参URL为：",url)
	fmt.Println("返回结果为",resp,err)
	utils.CatContext()
	resp,err = MyGetTemp(client,url)
	return
}




func MyGetTemp(client *http.Client,url string) (resp *http.Response, err error) {
	for i:=0;i<100;i++ {

	}
	return
}
