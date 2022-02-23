package bufioWriterWriteString

import (
	"bufio"
	"github.com/HXSecurity/DongTai-agent-go/global"
	"github.com/HXSecurity/DongTai-agent-go/utils"
)

func WriteString(b *bufio.Writer, v string) (n int, err error) {
	if utils.IsHook("net/http.extraHeader.Write", 6) {
		value, ok := global.ResponseHeaderMap.Load(utils.CatGoroutineID())
		if ok {
			str := utils.StringAdd(value.(string), string(v))
			global.ResponseHeaderMap.Store(utils.CatGoroutineID(), str)
		} else {
			global.ResponseHeaderMap.Store(utils.CatGoroutineID(), string(v))
		}
	}
	//request.FmtHookPool(request.PoolReq{
	//	Args:            request.Collect(argStr, s),
	//	Reqs:            request.Collect(b.String()),
	//	NeedHook:        request.Collect(argStr, s),
	//	NeedCatch:       request.Collect(b.String()),
	//	Source:          false,
	//	OriginClassName: "bytes.(*Buffer)",
	//	MethodName:      "WriteString",
	//	ClassName:       "bytes.(*Buffer)",
	//})
	return WriteStringT(b, v)
}

func WriteStringT(b *bufio.Writer, v string) (n int, err error) {
	return n, err
}
