package bufferWriteString

import (
	"bytes"
	"go-agent/model/request"
	"go-agent/utils"
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
