package urlURLQuery

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"net/url"
)

func Query(URL *url.URL) url.Values {
	values := QueryT(URL)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(),
		Reqs:            request.Collect(values),
		Source:          true,
		OriginClassName: "url.(*URL)",
		MethodName:      "Query",
		ClassName:       "url.(*URL)",
	})
	return values
}

func QueryT(URL *url.URL) url.Values {
	return url.Values{}
}
