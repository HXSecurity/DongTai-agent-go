package httpRequestFormValue

import (
	"go-agent/global"
	"go-agent/model/request"
	"go-agent/utils"
	"net/http"
)

func FormValue(req *http.Request, key string) string {
	signature, callerClass, callerMethod, callerLineNumber := utils.FmtStack()
	id := utils.CatGoroutineID()
	r := FormValueT(req, key)
	var sourceHash global.HashKeys = []string{utils.GetSource(key)}
	var targetHash global.HashKeys = []string{utils.GetSource(r)}
	var pool = request.Pool{
		Source:           true,
		Interfaces:       []interface{}{},
		SourceValues:     key,
		SourceHash:       sourceHash,
		TargetValues:     r,
		TargetHash:       targetHash,
		Signature:        signature,
		OriginClassName:  "http.(*Request)",
		MethodName:       "FormValue",
		ClassName:        "http.(*Request)",
		CallerLineNumber: callerLineNumber,
		CallerClass:      callerClass,
		CallerMethod:     callerMethod,
		RetClassName:     "string",
		Args:             key,
	}
	poolTree := request.PoolTree{
		Begin:       true,
		Pool:        &pool,
		Children:    []*request.PoolTree{},
		GoroutineID: id,
	}
	global.PoolTreeMap[&targetHash] = &poolTree
	return r
}

func FormValueT(r *http.Request, key string) string {
	return key
}
