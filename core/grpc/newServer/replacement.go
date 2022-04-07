package newServer

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/api"
	"github.com/HXSecurity/DongTai-agent-go/global"
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func NewServer(opt ...grpc.ServerOption) *grpc.Server {
	opt = append(opt, grpc.UnaryInterceptor(interceptor))
	return NewServerT(opt...)
}

// interceptor 一元拦截器
func interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	id := utils.CatGoroutineID()
	request.FmtHookPool(request.PoolReq{
		Reqs:            request.Collect(req),
		Source:          true,
		OriginClassName: "grpc",
		MethodName:      "NewServer",
		ClassName:       "grpc",
	})
	// 获取metadata
	res, err := handler(ctx, req)
	go func() {
		worker, _ := utils.NewWorker(global.AgentId)
		onlyKey := int(worker.GetId())
		header := base64.StdEncoding.EncodeToString([]byte(md.Get("headers")[0]))
		HookGroup := &request.UploadReq{
			Type:     36,
			InvokeId: onlyKey,
			Detail: request.Detail{
				AgentId: global.AgentId,
				Function: request.Function{
					Method:        "RPC",
					Url:           md.Get("requestURL")[0],
					Uri:           md.Get("requestURI")[0],
					Protocol:      md.Get("protocol")[0],
					ClientIp:      "",
					Language:      "GO",
					Scheme:        "GRPC",
					ReplayRequest: false,
					ReqHeader:     header,
					ReqBody:       "",
					QueryString:   "",
					Pool:          []request.Pool{},
					TraceId:       md.Get("dt-traceid")[0],
				},
			},
		}

		goroutineIDs := make(map[string]bool)
		global.PoolTreeMap.Range(func(key, value interface{}) bool {
			if value.(*request.PoolTree).IsThisBegin(id) {
				global.PoolTreeMap.Delete(key)
				value.(*request.PoolTree).FMT(&HookGroup.Detail.Function.Pool, worker, goroutineIDs, HookGroup.Detail.Function.TraceId)
				return false
			}
			return true
		})

		fmt.Println(HookGroup.Detail.Function.Url)
		api.ReportUpload(*HookGroup)
		request.RunMapGCbYGoroutineID(goroutineIDs)
	}()
	return res, err
}

func NewServerT(opt ...grpc.ServerOption) *grpc.Server {
	return nil
}
