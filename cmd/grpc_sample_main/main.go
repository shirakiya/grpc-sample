package main

import (
	"context"
	"log"
	"net"

	"github.com/shirakiya/grpc-sample/go/samplev1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type sampleServer struct {
	samplev1.UnimplementedSampleAPIServer
}

func (s *sampleServer) GetFoo(
	ctx context.Context,
	req *samplev1.FooRequest,
) (*samplev1.FooResponse, error) {
	log.Println("Requested GetFoo")

	if req.Foo == "" {
		return nil, status.Error(codes.InvalidArgument, "Foo must be set")
	}

	return &samplev1.FooResponse{
		Bar: "success!",
	}, nil
}

func newServer() *sampleServer {
	return &sampleServer{}
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	samplev1.RegisterSampleAPIServer(grpcServer, newServer())

	log.Println("Start server")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
