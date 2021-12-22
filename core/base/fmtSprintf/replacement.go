package fmtSprintf

import (
	"go-agent/model/request"
	"go-agent/utils"
)

func Sprintf(format string, a ...interface{}) string {
	s := SprintfT(format, a...)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(format, a),
		Reqs:            utils.Collect(s),
		Source:          false,
		OriginClassName: "fmt",
		MethodName:      "Sprintf",
		ClassName:       "fmt",
	})
	return s
}

func SprintfT(format string, a ...interface{}) string {
	return ""
}
