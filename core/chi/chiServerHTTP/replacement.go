/*
 *  @author: Breaker
 *  @date: 2022/3/28 6:29 PM
 */

package chiServerHTTP

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"net/http"
	"reflect"
	"strings"

	"github.com/HXSecurity/DongTai-agent-go/api"
	"github.com/HXSecurity/DongTai-agent-go/global"
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"github.com/go-chi/chi"
)

func MyServer(server *chi.Mux, w http.ResponseWriter, r *http.Request) {
	MyServerTemp(server, w, r)
	id := utils.CatGoroutineID()
	go func() {
		t := reflect.ValueOf(r.Body)
		var headerBase string
		body := ""
		for k, v := range r.Header {
			headerBase += k + ": " + strings.Join(v, ",") + "\n"
		}
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
		worker, _ := utils.NewWorker(global.AgentId)
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
				},
			},
		}
		var resBody string
		res, ok := global.ResponseMap.Load(id)
		if ok {
			global.ResponseMap.Delete(id)
			resBody = res.(string)
		}

		var resH string
		value2, ok1 := global.ResponseHeaderMap.Load(id)
		if ok1 {
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
				value.(*request.PoolTree).FMT(&HookGroup.Detail.Function.Pool, worker, goroutineIDs)
				return false
			}
			return true
		})
		api.ReportUpload(*HookGroup)
		request.RunMapGCbYGoroutineID(goroutineIDs)
	}()
	return
}

func MyServerTemp(server *chi.Mux, w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 100; i++ {

	}
	return
}
