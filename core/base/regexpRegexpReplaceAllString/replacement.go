package regexpRegexpReplaceAllString

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"regexp"
)

func ReplaceAllString(re *regexp.Regexp, src, repl string) string {
	s := ReplaceAllStringT(re, src, repl)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(src, repl),
		Reqs:            request.Collect(s),
		Source:          false,
		OriginClassName: "regexp.(*Regexp)",
		MethodName:      "ReplaceAllString",
		ClassName:       "regexp.(*Regexp)",
	})
	return s
}

func ReplaceAllStringT(re *regexp.Regexp, src, repl string) string {
	return ""
}
