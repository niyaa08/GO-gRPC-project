package main

import (
	"context"
	pb "github.com/niyaa08/grpc-project/proto"
)

func (s *helloServer) SayHi(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
