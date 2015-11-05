package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "github.com/ppg/grpc-intro/proto"
)

func main() {
	v1, v2 := 12, 23

	// Dial the server insecurely
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln("fail to dial:", err)
	}
	defer conn.Close()

	// Get a client using the connection
	client := pb.NewTestClient(conn)
	resp, err := client.Add(context.Background(), &pb.NumericRequest{V1: int32(v1), V2: int32(v2)})
	if err != nil {
		grpclog.Fatalln("failed server call:", err)
	}
	grpclog.Println("resp:", resp)
}
