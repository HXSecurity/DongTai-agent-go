package stringsBuilderWriteString

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"strings"
)

func WriteString(b *strings.Builder, s string) (n int, err error) {
	argStr := b.String()
	n, err = WriteStringT(b, s)
	var WriteMap = make(map[string]bool)
	WriteMap["strings.Join"] = true
	WriteMap["strings.Replace"] = true
	WriteMap["net/url.Values.Encode"] = true
	WriteMap["net/url.(*URL).String"] = true
	if !WriteMap[utils.LoadFunc(2)] {
		request.FmtHookPool(request.PoolReq{
			Args:            request.Collect(argStr, s),
			Reqs:            request.Collect(b.String()),
			NeedHook:        request.Collect(argStr, s),
			NeedCatch:       request.Collect(b.String()),
			Source:          false,
			OriginClassName: "strings.(*Builder)",
			MethodName:      "WriteString",
			ClassName:       "strings.(*Builder)",
		})
	}
	return n, err
}

func WriteStringT(b *strings.Builder, s string) (n int, err error) {
	return n, err
}
