package main

import (
	"flag"
	"log"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	args := flag.Args()
	v1, _ := strconv.Atoi(args[0])
	v2, _ := strconv.Atoi(args[1])

	// Accumulate gRPC options
	var opts []grpc.DialOption

	// Connect via insecure
	opts = append(opts, grpc.WithInsecure())

	// Dial the server
	conn, err := grpc.Dial("127.0.0.1:1234", opts...)
	if err != nil {
		log.Fatalln("fail to dial:", err)
	}
	defer conn.Close()

	// Get a client using the connection
	client := NewTestClient(conn)
	resp, err := client.Add(context.Background(), &NumericRequest{V1: int32(v1), V2: int32(v2)})
	if err != nil {
		log.Fatalln("failed server call:", err)
	}
	log.Println("resp:", resp)
}
