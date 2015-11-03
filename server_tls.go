package main

import (
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

func main() {
	// Create a TCP listener like you normally would
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	// Create an instance of a struct that implements the TestServer interface (i.e. Add)
	service := &testService{}

	// Accumulate gRPC options
	var opts []grpc.ServerOption

	// Serve via TLS
	creds, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	if err != nil {
		grpclog.Fatalf("Failed to generate credentials %v", err)
	}
	opts = append(opts, grpc.Creds(creds))

	// Create a gRPC server and attach our service to it
	server := grpc.NewServer(opts...)
	RegisterTestServer(server, service)

	// Start handling requests
	grpclog.Printf("Starting to listen on %s", listener.Addr())
	err = server.Serve(listener)
	if err != nil {
		grpclog.Fatalf("failed to serve: %v", err)
	}
}

type testService struct {
}

func (s *testService) Add(ctx context.Context, req *NumericRequest) (*NumericResponse, error) {
	grpclog.Printf("v1=%d, v2=%s", req.V1, req.V2)
	return &NumericResponse{R: req.V1 + req.V2}, nil
}
