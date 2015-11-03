package main

import (
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// START SERVICE IMPLEMENTATION OMIT
type testService struct {
}

func (s *testService) Add(ctx context.Context, req *NumericRequest) (*NumericResponse, error) {
	grpclog.Printf("v1=%d, v2=%d", req.V1, req.V2)
	return &NumericResponse{R: req.V1 + req.V2}, nil
}

// END SERVICE IMPLEMENTATION OMIT

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	// Create a gRPC server and attach our service to it
	server := grpc.NewServer()
	RegisterTestServer(server, &testService{})

	// Start handling requests
	grpclog.Printf("Starting to listen on %s", listener.Addr())
	if err = server.Serve(listener); err != nil {
		grpclog.Fatalf("failed to serve: %v", err)
	}
}
