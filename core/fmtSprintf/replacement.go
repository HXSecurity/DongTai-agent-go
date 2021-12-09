package fmtSprintf

import (
	"go-agent/global"
	"go-agent/model/request"
	"go-agent/utils"
)

func Sprintf(format string, a ...interface{}) string {
	signature, callerClass, callerMethod, callerLineNumber := utils.FmtStack()
	var sourceHash global.HashKeys = []string{}
	var SourceValues string = ""
	for _, v := range a {
		switch v.(type) {
		case string:
			sourceHash = append(sourceHash, utils.GetSource(v))
			SourceValues += utils.StringAdd(v.(string), " ")
		}
	}
	s := SprintfT(format, a...)

	var targetHash global.HashKeys = []string{utils.GetSource(s)}

	var pool = request.Pool{
		Source:           false,
		Interfaces:       []interface{}{},
		SourceValues:     SourceValues,
		SourceHash:       sourceHash,
		TargetValues:     s,
		TargetHash:       targetHash,
		Signature:        signature,
		OriginClassName:  "fmt",
		MethodName:       "Sprintf",
		ClassName:        "fmt",
		CallerLineNumber: callerLineNumber,
		CallerClass:      callerClass,
		CallerMethod:     callerMethod,
		RetClassName:     "string",
		Args:             utils.Strval(a),
	}

	poolTree := request.PoolTree{
		Begin:       false,
		Pool:        &pool,
		Children:    []*request.PoolTree{},
		GoroutineID: utils.CatGoroutineID(),
	}
	global.PoolTreeMap.Range(func(key, value interface{}) bool {
		if key.(*global.HashKeys).Some(sourceHash) {
			value.(*request.PoolTree).Children = append(value.(*request.PoolTree).Children, &poolTree)
			return false
		}
		return true
	})
	global.PoolTreeMap.Store(&targetHash, &poolTree)
	return s
}

func SprintfT(format string, a ...interface{}) string {
	return ""
}
