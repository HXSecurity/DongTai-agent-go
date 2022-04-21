package httpRouter

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/api"
	"github.com/HXSecurity/DongTai-agent-go/global"
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func MyHttpRouterServer(server *httprouter.Router, w http.ResponseWriter, r *http.Request) {
	worker, err := utils.NewWorker(global.AgentId)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(r.Host),
		Reqs:            request.Collect(r.Host),
		Source:          true,
		OriginClassName: "httprouter.(*Router)",
		MethodName:      "ServeHTTP",
		ClassName:       "httprouter",
	})

	MyHttpRouterServerTemp(server, w, r)
	id := utils.CatGoroutineID()
	go func() {
		t := reflect.ValueOf(r.Body)
		var headerBase string
		body := ""
		for k, v := range r.Header {
			headerBase += k + ": " + strings.Join(v, ",") + "\n"
		}
		TraceId := global.TraceId + "-" + strconv.Itoa(int(worker.GetId()))
		global.TargetTraceId = TraceId
		tranceID := TraceId + "." + strconv.Itoa(global.AgentId) + ".0.0.0"
		headerBase += "dt-traceid:" + tranceID
		if t.Kind() == reflect.Ptr {
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
		header := base64.StdEncoding.EncodeToString([]byte(headerBase))
		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}
		if err != nil {
			fmt.Println(err)
		}
		onlyKey := int(worker.GetId())
		HookGroup := &request.UploadReq{
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
					TraceId:       tranceID,
				},
			},
		}
		var resBody string
		var resH string
		res, ok := global.ResponseMap.Load(id)
		if ok {
			global.ResponseMap.Delete(id)
			resBody = res.(string)
		}
		value2, ok2 := global.ResponseHeaderMap.Load(id)
		if ok2 {
			global.ResponseHeaderMap.Delete(id)
			resH = value2.(string)
		}
		for k, v := range w.Header() {
			resH += k + ": " + strings.Join(v, ",") + "\n"
		}
		resHeader := base64.StdEncoding.EncodeToString([]byte(resH))
		HookGroup.Detail.ResHeader = resHeader
		HookGroup.Detail.ResBody = resBody
		goroutineIDs := make(map[string]bool)
		global.PoolTreeMap.Range(func(key, value interface{}) bool {
			if value.(*request.PoolTree).IsThisBegin(id) {
				global.PoolTreeMap.Delete(key)
				value.(*request.PoolTree).FMT(&HookGroup.Detail.Function.Pool, worker, goroutineIDs, HookGroup.Detail.Function.TraceId)
				return true
			}
			return true
		})
		api.ReportUpload(*HookGroup)
		request.RunMapGCbYGoroutineID(goroutineIDs)
	}()
	return
}

func MyHttpRouterServerTemp(server *httprouter.Router, w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 100; i++ {

	}
	return
}
