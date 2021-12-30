package bufferWriteString

import (
	"bytes"
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
)

func WriteString(b *bytes.Buffer, s string) (n int, err error) {
	argStr := b.String()
	n, err = WriteStringT(b, s)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(argStr, s),
		Reqs:            utils.Collect(b.String()),
		NeedHook:        utils.Collect(argStr, s),
		NeedCatch:       utils.Collect(b.String()),
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
