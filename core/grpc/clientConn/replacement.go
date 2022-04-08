package clientConn

import (
	"context"
	"github.com/HXSecurity/DongTai-agent-go/global"
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strconv"
	"strings"
)

func Invoke(cl *grpc.ClientConn, ctx context.Context, method string, args, reply interface{}, opts ...grpc.CallOption) error {
	outmd, _ := metadata.FromIncomingContext(ctx)
	worker, _ := utils.NewWorker(global.AgentId)
	var tranceid string
	if len(outmd.Get("dt-traceid")) > 0 {
		tranceid = outmd.Get("dt-traceid")[0]
	}

	if tranceid == "" {
		tranceid = global.TraceId + "." + strconv.Itoa(global.AgentId) + ".0.0." + strconv.Itoa(int(worker.GetId()))
	} else {
		four := strconv.Itoa(int(worker.GetId()))
		tranceids := strings.Split(tranceid, ".")
		tranceids[1] = strconv.Itoa(global.AgentId)
		num, _ := strconv.Atoi(tranceids[3])
		tranceids[3] = strconv.Itoa(num + 1)
		tranceids[4] = four
		newId := ""
		for i := 0; i < len(tranceids); i++ {
			if i == 4 {
				newId += tranceids[i]
			} else {
				newId += tranceids[i] + "."
			}
		}
		tranceid = newId
	}
	md := metadata.Pairs("dt-traceid", tranceid,
		"protocol", "ProtoBuf",
		"requestURL", cl.Target()+method,
		"requestURI", method,
		"headers", "traceid:"+tranceid,
	)
	ctx = metadata.NewOutgoingContext(ctx, md)
	err := InvokeT(cl, ctx, method, args, reply, opts...)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(args),
		Reqs:            request.Collect(reply),
		Source:          false,
		OriginClassName: "grpc.(*ClientConn)",
		MethodName:      "Invoke",
		ClassName:       "grpc.(*ClientConn)",
	})
	return err
}
func InvokeT(cl *grpc.ClientConn, ctx context.Context, method string, args, reply interface{}, opts ...grpc.CallOption) error {
	return nil
}
