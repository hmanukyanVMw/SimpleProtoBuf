// server.go
package main

import (
	"context"
	"log"
	"net"

	pb "github.com/hmanukyanVMw/SimpleProtoBuf/gen/go/sso"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) MyMethod(ctx context.Context, in *pb.MyRequest) (*pb.MyResponse, error) {
	return &pb.MyResponse{Result: "Hello " + in.Data}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
