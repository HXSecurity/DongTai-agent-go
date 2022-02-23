package fmtSprintf

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
)

func Sprintf(format string, a ...interface{}) string {
	s := SprintfT(format, a...)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(format, a),
		Reqs:            request.Collect(s),
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
