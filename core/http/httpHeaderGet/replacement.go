package httpHeaderGet

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"net/http"
	"strings"
)

func Get(header http.Header, key string) string {
	values := GetT(header, key)
	skipMap := make(map[string]bool)
	skipMap["net/http.(*persistConn).roundTrip"] = true
	skipMap["net/http.(*chunkWriter).writeHeader"] = true
	skipMap["net/http.(*Client).makeHeadersCopier"] = true
	skipMap["net/http.(*persistConn).readLoop"] = true
	skipMap["net/http.(*Request).requiresHTTP1"] = true
	skipMap["net/http.isProtocolSwitchHeader"] = true
	if strings.Index(utils.LoadFunc(2), "github.com/parnurzeal/gorequest") > -1 {
		return values
	}
	if skipMap[utils.LoadFunc(2)] {
		return values
	}
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(key),
		Reqs:            request.Collect(values),
		Source:          true,
		OriginClassName: "http.(Header)",
		MethodName:      "Get",
		ClassName:       "http.(Header)",
	})
	return values
}

func GetT(header http.Header, key string) string {
	return ""
}
