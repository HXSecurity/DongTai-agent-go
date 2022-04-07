package newServer

import (
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"google.golang.org/grpc"
)

func init() {
	model.HookMap["grpcNewServer"] = new(GrpcNewServer)
}

type GrpcNewServer struct {
}

func (h *GrpcNewServer) Hook() {
	err := gohook.Hook(grpc.NewServer, NewServer, NewServerT)
	if err != nil {
		fmt.Println(err, "GrpcNewServer")
	} else {
		fmt.Println("GrpcNewServer")
	}
}

func (h *GrpcNewServer) UnHook() {
	gohook.UnHook(grpc.NewServer)
}
