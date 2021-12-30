package fmtSprintf

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
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
