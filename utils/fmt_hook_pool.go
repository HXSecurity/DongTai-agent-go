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

	var NeedHook []interface{}
	RangeSource(p.Args, &NeedHook)
	for _, v := range NeedHook {
		switch v.(type) {
		case string:
			sourceHash = append(sourceHash, GetSource(v))
			SourceValues = StringAdd(SourceValues, v.(string), " ")
		}
	}
	var targetHash global.HashKeys
	var targetValues string = ""
	var RetClassNames string = ""
	var NeedCatch []interface{}

	RangeSource(p.Reqs, &NeedCatch)
	for _, v := range p.Reqs {
		if reflect.ValueOf(v).IsValid() {
			RetClassNames = StringAdd(RetClassNames, reflect.ValueOf(v).Type().String(), " ")
		}
	}
	for _, v := range NeedCatch {
		switch v.(type) {
		case string:
			targetHash = append(targetHash, GetSource(v))
			StringAdd(targetValues, v.(string), " ")
		}
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
		Args:             Strval(p.Args),
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
