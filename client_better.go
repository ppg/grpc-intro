package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "github.com/ppg/grpc-intro/proto"
)

func main() {
	values := []int32{1, 1, 2, 3, 5, 8}
	testType := pb.TestType_TYPE_1

	// Dial the server
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln("fail to dial:", err)
	}
	defer conn.Close()

	// Get a client using the connection
	client := pb.NewBetterTestClient(conn)
	resp, err := client.Add(context.Background(), &pb.BetterNumericRequest{Type: testType, Values: values})
	if err != nil {
		grpclog.Fatalln("failed server call:", err)
	}
	grpclog.Println("resp:", resp)
}
