package unaryInterceptor

import (
	"google.golang.org/grpc"
)

const (
	TraceId = iota
	AgentId
	RoutineId
	NextKey
	OnlyKey
)

var UnaryServerInterceptors []grpc.UnaryServerInterceptor

func UnaryInterceptor(i grpc.UnaryServerInterceptor) grpc.ServerOption {

	UnaryServerInterceptors = append(UnaryServerInterceptors, i)
	return nil
}

func UnaryInterceptorT(i grpc.UnaryServerInterceptor) grpc.ServerOption {
	return nil
}
