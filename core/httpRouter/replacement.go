package httpRouter

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-agent/api"
	"go-agent/global"
	"go-agent/model/request"
	"go-agent/utils"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

func MyHttpRouterServer(server *httprouter.Router, w http.ResponseWriter, r *http.Request) {
	id := utils.CatGoroutineID()
	//fmt.Println(r.Body,"body")
	t := reflect.ValueOf(r.Body)
	b, err := json.Marshal(r.Header)
	if err != nil {
		fmt.Println(err)
		return
	}
	header := base64.StdEncoding.EncodeToString(b)
	body := ""
	if t.Type().String() != "http.noBody" {
		buf := t.
			Elem().
			FieldByName("src").
			Elem().Elem().
			FieldByName("R").
			Elem().Elem().
			FieldByName("buf").Bytes()
		buf = buf[:bytes.IndexByte(buf, 0)]
		reader := bufio.NewReader(bytes.NewReader(buf))
		var reqArr []string
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				break
			}
			reqArr = append(reqArr, string(line))
		}
		body = reqArr[len(reqArr)-1]
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	onlyKey, err := strconv.Atoi(strconv.Itoa(global.AgentId) + id + strconv.Itoa(int(time.Now().Unix())))
	if err != nil {
		return
	}
	global.HookGroup[id] = &request.UploadReq{
		Type:     36,
		InvokeId: onlyKey,
		Detail: request.Detail{
			AgentId: global.AgentId,
			Function: request.Function{
				Method:        r.Method,
				Url:           scheme + "://" + r.Host + r.RequestURI,
				Uri:           r.URL.Path,
				Protocol:      r.Proto,
				ClientIp:      r.RemoteAddr,
				Language:      "GO",
				ReplayRequest: false,
				ReqHeader:     header,
				ReqBody:       body,
				QueryString:   r.URL.RawQuery,
				Pool:          []request.Pool{},
			},
		},
	}
	//fmt.Println("请求到了\n")
	//fmt.Println("method", r.Method)
	//fmt.Println("scheme", scheme)
	//fmt.Println("secure", r.TLS)
	//fmt.Println("url", scheme+"://"+r.Host+r.RequestURI)
	//fmt.Println("uri", r.URL.Path)
	//fmt.Println("protocol", r.Proto)
	//fmt.Println("clientId", r.RemoteAddr)
	//fmt.Println("language", "GO")
	//fmt.Println("replayRequest", false)
	//fmt.Println("header", baseStr)
	//fmt.Println("reqBody", string(s))
	//fmt.Println("queryString", r.URL.RawQuery)
	MyHttpRouterServerTemp(server, w, r)
	resH, err := json.Marshal(w.Header())
	if err != nil {
		fmt.Println(err)
		return
	}
	resBody := string(reflect.ValueOf(w).Elem().FieldByName("w").Elem().FieldByName("buf").Bytes())
	resHeader := base64.StdEncoding.EncodeToString(resH)
	global.HookGroup[id].Detail.ResHeader = resHeader
	global.HookGroup[id].Detail.ResBody = resBody
	goroutineIDs := make(map[string]bool)
	global.PoolTreeMap.Range(func(key, value interface{}) bool {
		if value.(*request.PoolTree).IsThisBegin(id) {
			onlyKey += 1
			value.(*request.PoolTree).FMT(&global.HookGroup[id].Detail.Function.Pool, onlyKey, goroutineIDs)
			return true
		}
		return true
	})
	api.ReportUpload(*global.HookGroup[id])
	utils.RunMapGCbYGoroutineID(goroutineIDs)
	return
}

func MyHttpRouterServerTemp(server *httprouter.Router, w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 100; i++ {

	}
	return
}
