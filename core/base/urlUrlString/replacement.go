package urlUrlString

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"net/url"
)

func String(url *url.URL) string {
	s := StringT(url)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(url),
		Reqs:            request.Collect(s),
		Source:          false,
		OriginClassName: "url.(*URL)",
		MethodName:      "String",
		ClassName:       "url.(*URL)",
	})

	return s
}

func StringT(url *url.URL) string {
	return ""
}
