package bufioWriterWrite

import (
	"bufio"
	"github.com/HXSecurity/DongTai-agent-go/global"
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"reflect"
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

		var u uintptr
		va := reflect.ValueOf(v)
		u = va.Pointer()

		utils.FmtHookPool(request.PoolReq{
			Args:            utils.Collect(b, v),
			Reqs:            utils.Collect(n, err),
			NeedHook:        utils.Collect(u),
			Source:          false,
			OriginClassName: "bufio.(*Buffer)",
			MethodName:      "Write",
			ClassName:       "bufio.(*Buffer)",
		})
	}
	if utils.IsHook("net/http.extraHeader.Write", 6) {
		value, ok := global.ResponseHeaderMap.Load(utils.CatGoroutineID())
		if ok {
			str := utils.StringAdd(value.(string), string(v))
			global.ResponseHeaderMap.Store(utils.CatGoroutineID(), str)
		} else {
			global.ResponseHeaderMap.Store(utils.CatGoroutineID(), string(v))
		}
	}

	return WriteT(b, v)
}

func WriteT(b *bufio.Writer, v []byte) (n int, err error) {
	return n, err
}
