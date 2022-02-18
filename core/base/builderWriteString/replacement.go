package builderWriteString

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"strings"
)

func WriteString(b *strings.Builder, s string) (n int, err error) {
	argStr := b.String()
	n, err = WriteStringT(b, s)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(argStr, s),
		Reqs:            request.Collect(b.String()),
		Source:          false,
		OriginClassName: "strings.(*Builder)",
		MethodName:      "WriteString",
		ClassName:       "strings.(*Builder)",
	})
	return n, err
}

func WriteStringT(b *strings.Builder, s string) (n int, err error) {
	return n, err
}
