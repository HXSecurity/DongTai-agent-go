package bufferWriteString

import (
	"bytes"
	"github.com/HXSecurity/DongTai-agent-go/model/request"
)

func WriteString(b *bytes.Buffer, s string) (n int, err error) {
	argStr := b.String()
	n, err = WriteStringT(b, s)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(argStr, s),
		Reqs:            request.Collect(b.String()),
		NeedHook:        request.Collect(argStr, s),
		NeedCatch:       request.Collect(b.String()),
		Source:          false,
		OriginClassName: "bytes.(*Buffer)",
		MethodName:      "WriteString",
		ClassName:       "bytes.(*Buffer)",
	})
	return n, err
}

func WriteStringT(b *bytes.Buffer, s string) (n int, err error) {
	return n, err
}
