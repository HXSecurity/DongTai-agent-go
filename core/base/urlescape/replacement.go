package urlescape

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	_ "unsafe"
)

type encoding int

const (
	encodePath encoding = 1 + iota
	encodePathSegment
	encodeHost
	encodeZone
	encodeUserPassword
	encodeQueryComponent
	encodeFragment
)

//go:linkname escape net/url.escape
func escape(s string, mode encoding) string

func escapeR(s string, mode encoding) string {
	b := escapeT(s, mode)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(s, mode),
		Reqs:            utils.Collect(b),
		Source:          false,
		OriginClassName: "url",
		MethodName:      "escape",
		ClassName:       "url",
	})
	return b
}

func escapeT(s string, mode encoding) string {
	return ""
}
