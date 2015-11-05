package main

import (
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	pb "github.com/ppg/grpc-intro/proto"
)

type testService struct {
}

func (s *testService) Add(ctx context.Context, req *pb.NumericRequest) (*pb.NumericResponse, error) {
	grpclog.Printf("v1=%d, v2=%d", req.V1, req.V2)
	return &pb.NumericResponse{R: req.V1 + req.V2}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	// Serve via TLS
	creds, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	if err != nil {
		grpclog.Fatalf("Failed to generate credentials %v", err)
	}
	opts := []grpc.ServerOption{grpc.Creds(creds)}

	// Create a gRPC server and attach our service to it
	server := grpc.NewServer(opts...)
	pb.RegisterTestServer(server, &testService{})

	// Start handling requests
	grpclog.Printf("Starting to listen on %s", listener.Addr())
	if err = server.Serve(listener); err != nil {
		grpclog.Fatalf("failed to serve: %v", err)
	}
}
