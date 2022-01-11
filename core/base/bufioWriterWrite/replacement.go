package bufioWriterWrite

import (
	"bufio"
	"github.com/HXSecurity/DongTai-agent-go/global"
	"github.com/HXSecurity/DongTai-agent-go/utils"
)

func Write(b *bufio.Writer, v []byte) (n int, err error) {
	if utils.IsHook("net/http.(*response).write", 6) {
		value, ok := global.ResponseMap.Load(utils.CatGoroutineID())
		if ok {
			str := utils.StringAdd(value.(string), string(v))
			global.ResponseMap.Store(utils.CatGoroutineID(), str)
		} else {
			global.ResponseMap.Store(utils.CatGoroutineID(), string(v))
		}
	}
	//utils.FmtHookPool(request.PoolReq{
	//	Args:            utils.Collect(argStr, s),
	//	Reqs:            utils.Collect(b.String()),
	//	NeedHook:        utils.Collect(argStr, s),
	//	NeedCatch:       utils.Collect(b.String()),
	//	Source:          false,
	//	OriginClassName: "bytes.(*Buffer)",
	//	MethodName:      "WriteString",
	//	ClassName:       "bytes.(*Buffer)",
	//})
	return WriteT(b, v)
}

func WriteT(b *bufio.Writer, v []byte) (n int, err error) {
	return n, err
}
