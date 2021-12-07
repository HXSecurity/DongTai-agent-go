package httpRequestFormValue

import (
	"go-agent/global"
	"go-agent/model/request"
	"go-agent/utils"
	"net/http"
)

func FormValue(req *http.Request, key string) string {
	r := FormValueT(req, key)
	var sourceHash global.HashKeys = []string{utils.GetSource(key)}
	var targetHash global.HashKeys = []string{utils.GetSource(r)}
	var pool = request.Pool{
		SourceValues: key,
		SourceHash:   sourceHash,
		TargetHash:   targetHash,
	}
	poolTree := request.PoolTree{
		Begin:       true,
		Pool:        &pool,
		Children:    []*request.PoolTree{},
		GoroutineID: utils.CatGoroutineID(),
	}
	global.PoolTreeMap[&targetHash] = &poolTree
	return r
}

func FormValueT(r *http.Request, key string) string {
	return key
}
