package grpc

import (
	"context"
	"fmt"

	"github.com/eavehh/marketpl.microserv/internal/promotion/grpc/pb"
)

type Greeter_service struct {
	pb.UnimplementedGreeterServer
}

func New_greeter_service() *Greeter_service {
	return &Greeter_service{}
}

func (s *Greeter_service) SayHello(_ context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Msg: fmt.Sprintf("Hello %s", req.GetName()),
	}, nil
}
