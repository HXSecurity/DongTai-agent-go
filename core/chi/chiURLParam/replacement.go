/*
 *  @author: Breaker
 *  @date: 2022/3/29 10:26 AM
 */

package chiURLParam

import (
	"net/http"

	"github.com/HXSecurity/DongTai-agent-go/model/request"
)

func chiURLParam(r *http.Request, key string) string {
	str := chiURLParamTemp(r, key)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(key),
		Reqs:            request.Collect(str),
		Source:          true,
		OriginClassName: "chi",
		MethodName:      "URLParam",
		ClassName:       "chi",
	})
	return str
}

func chiURLParamTemp(r *http.Request, key string) string {
	return ""
}
