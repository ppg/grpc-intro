package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	pb "./proto"
)

func main() {
	v1, v2 := int32(12), int32(23)

	// Get TLS creds
	creds, err := credentials.NewClientTLSFromFile("server.crt", "grpc.bouldergo.com")
	if err != nil {
		grpclog.Fatalf("Failed to generate credentials %v", err)
	}

	// Dial the server securely
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithTransportCredentials(creds))
	if err != nil {
		grpclog.Fatalln("fail to dial:", err)
	}
	defer conn.Close()

	// Get a client using the connection
	client := pb.NewTestClient(conn)
	resp, err := client.Add(context.Background(), &pb.NumericRequest{V1: v1, V2: v2})
	if err != nil {
		grpclog.Fatalln("failed server call:", err)
	}
	grpclog.Println("resp:", resp)
}
