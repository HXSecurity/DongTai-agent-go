package builderWriteString

import (
	"go-agent/model/request"
	"go-agent/utils"
	"strings"
)

func WriteString(b *strings.Builder, s string) (n int, err error) {
	argStr := b.String()
	n, err = WriteStringT(b, s)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(argStr, s),
		Reqs:            utils.Collect(b.String()),
		Source:          false,
		OriginClassName: "bytes.(*Buffer)",
		MethodName:      "WriteString",
		ClassName:       "bytes.(*Buffer)",
	})
	return n, err
}

func WriteStringT(b *strings.Builder, s string) (n int, err error) {
	return n, err
}
