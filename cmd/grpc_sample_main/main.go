package main

import (
	"context"
	"log"
	"net"
	"strings"

	"github.com/shirakiya/grpc-sample/go/samplev1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
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

func getAuthInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		service, method := parseMethod(info.FullMethod)

		opts := req.(proto.Message).ProtoReflect().Descriptor().ParentFile().Services().ByName(protoreflect.Name(service)).Methods().ByName(protoreflect.Name(method)).Options().(*descriptorpb.MethodOptions)
		auth := proto.GetExtension(opts, samplev1.E_Authentication).(*samplev1.Authentication)

		log.Printf("Required: %t", auth.Required)

		if auth.Required {
			log.Println("Required authentication")
		}

		return handler(ctx, req)
	}
}

func parseMethod(s string) (string, string) {
	// Expected below string.
	//
	// /sample.v1.SampleAPI/GetFoo
	//
	subs := strings.Split(s, "/")

	serviceSub := strings.Split(subs[1], ".")
	service := serviceSub[len(serviceSub)-1]

	method := subs[len(subs)-1]

	return service, method
}

func buildGrpcOptions() []grpc.ServerOption {
	interceptor := grpc.ChainUnaryInterceptor(
		getAuthInterceptor(),
	)

	return []grpc.ServerOption{
		interceptor,
	}
}

func newServer() *sampleServer {
	return &sampleServer{}
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcOptions := buildGrpcOptions()

	grpcServer := grpc.NewServer(grpcOptions...)
	samplev1.RegisterSampleAPIServer(grpcServer, newServer())

	log.Println("Start server")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
