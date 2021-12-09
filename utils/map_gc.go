package utils

import (
	"go-agent/global"
	"go-agent/model/request"
)

func RunMapGCbYGoroutineID(goroutineIDs map[string]bool) {

	global.PoolTreeMap.Range(func(key, value interface{}) bool {
		for id, _ := range goroutineIDs {
			if value.(*request.PoolTree).GoroutineID == id {
				global.PoolTreeMap.Delete(key)
			}
		}
		return true
	})
}
