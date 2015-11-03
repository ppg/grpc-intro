package main

import (
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	// Create a TCP listener like you normally would
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	// Create an instance of a struct that implements the TestServer interface (i.e. Add)
	service := &testBetterService{}

	// Accumulate gRPC options
	var opts []grpc.ServerOption

	// Create a gRPC server and attach our service to it
	server := grpc.NewServer(opts...)
	RegisterBetterTestServer(server, service)

	// Start handling requests
	grpclog.Printf("Starting to listen on %s", listener.Addr())
	err = server.Serve(listener)
	if err != nil {
		grpclog.Fatalf("failed to serve: %v", err)
	}
}

type testBetterService struct {
}

func (s *testBetterService) Add(ctx context.Context, req *BetterNumericRequest) (*BetterNumericResponse, error) {
	var sum int32
	for _, v := range req.Values {
		grpclog.Printf(" + %v", v)
		sum += v
	}
	prefix := ""
	switch req.Type {
	case TestType_TYPE_0:
		prefix = "This is type 0"
	case TestType_TYPE_1:
		prefix = "This is type 1"
	}
	return &BetterNumericResponse{Prefix: prefix, R: sum}, nil
}
