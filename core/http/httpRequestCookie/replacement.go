package httpRequestCookie

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"net/http"
)

func Cookie(req *http.Request, name string) (*http.Cookie, error) {
	cookie, err := CookieT(req, name)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(name),
		Reqs:            request.Collect(cookie, err),
		Source:          true,
		OriginClassName: "http.(*Request)",
		MethodName:      "Cookie",
		ClassName:       "http.(*Request)",
	})
	return cookie, err
}

func CookieT(req *http.Request, name string) (c *http.Cookie, e error) {
	return c, e
}
