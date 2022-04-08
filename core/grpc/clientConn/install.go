package clientConn

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"google.golang.org/grpc"
)

func init() {
	model.HookMap["grpcClientConn"] = new(GrpcClientConn)
}

type GrpcClientConn struct {
}

func (h *GrpcClientConn) Hook() {
	cl := &grpc.ClientConn{}
	err := gohook.HookMethod(cl, "Invoke", Invoke, InvokeT)
	if err != nil {
		fmt.Println(err, "GrpcClientConn")
	} else {
		fmt.Println("GrpcClientConn")
	}
}

func (h *GrpcClientConn) UnHook() {
	cl := &grpc.ClientConn{}
	gohook.UnHookMethod(cl, "Invoke")
}
