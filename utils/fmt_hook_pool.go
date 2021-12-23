package utils

import (
	"go-agent/global"
	"go-agent/model/request"
	"reflect"
)

func Collect(args ...interface{}) []interface{} {
	return args
}

func FmtHookPool(p request.PoolReq) {
	signature, callerClass, callerMethod, callerLineNumber := FmtStack()
	var sourceHash global.HashKeys
	var SourceValues string = ""
	if len(p.NeedHook) == 0 {
		RangeSource(p.Args, &p.NeedHook)
	}
	for _, v := range p.NeedHook {
		switch v.(type) {
		case string:
			sourceHash = append(sourceHash, GetSource(v))
			SourceValues = StringAdd(SourceValues, v.(string), " ")
		}
	}
	var targetHash global.HashKeys
	var targetValues string = ""
	var RetClassNames string = ""
	if len(p.NeedCatch) == 0 {
		RangeSource(p.Reqs, &p.NeedCatch)
	}
	for _, v := range p.Reqs {
		if reflect.ValueOf(v).IsValid() {
			RetClassNames = StringAdd(RetClassNames, reflect.ValueOf(v).Type().String(), " ")
		}
	}
	for _, v := range p.NeedCatch {
		switch v.(type) {
		case string:
			targetHash = append(targetHash, GetSource(v))
			targetValues = StringAdd(targetValues, v.(string), " ")
		}
	}
	var ArgsStr string
	if p.ArgsStr != "" {
		ArgsStr = p.ArgsStr
	} else {
		ArgsStr = Strval(p.Args)
	}
	var pool = request.Pool{
		Source:           p.Source,
		Interfaces:       []interface{}{},
		SourceValues:     SourceValues,
		SourceHash:       sourceHash,
		TargetValues:     targetValues,
		TargetHash:       targetHash,
		Signature:        signature,
		OriginClassName:  p.OriginClassName,
		MethodName:       p.MethodName,
		ClassName:        p.ClassName,
		CallerLineNumber: callerLineNumber,
		CallerClass:      callerClass,
		CallerMethod:     callerMethod,
		RetClassName:     RetClassNames,
		Args:             ArgsStr,
	}

	poolTree := request.PoolTree{
		Begin:       p.Source,
		Pool:        &pool,
		Children:    []*request.PoolTree{},
		GoroutineID: CatGoroutineID(),
	}
	if !p.Source {
		global.PoolTreeMap.Range(func(key, value interface{}) bool {
			if key.(*global.HashKeys).Some(sourceHash) {
				value.(*request.PoolTree).Children = append(value.(*request.PoolTree).Children, &poolTree)
				return false
			}
			return true
		})
	}
	global.PoolTreeMap.Store(&targetHash, &poolTree)
}
