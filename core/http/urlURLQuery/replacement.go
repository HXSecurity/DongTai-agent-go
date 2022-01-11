package urlURLQuery

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"net/url"
)

func Query(URL *url.URL) url.Values {
	values := QueryT(URL)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(),
		Reqs:            utils.Collect(values),
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
