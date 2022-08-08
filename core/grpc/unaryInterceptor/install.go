package unaryInterceptor

import (
	"fmt"

	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"google.golang.org/grpc"
)

func init() {
	model.HookMap["grpcUnaryInterceptor"] = new(GrpcUnaryInterceptor)
}

type GrpcUnaryInterceptor struct {
}

func (h *GrpcUnaryInterceptor) Hook() {
	err := gohook.Hook(grpc.UnaryInterceptor, UnaryInterceptor, nil)
	if err != nil {
		fmt.Println(err, "GrpcUnaryInterceptor")
	} else {
		fmt.Println("GrpcUnaryInterceptor")
	}
}

func (h *GrpcUnaryInterceptor) UnHook() {
	gohook.UnHook(grpc.UnaryInterceptor)
}
