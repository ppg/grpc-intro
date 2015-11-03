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
	var values []int32
	for _, arg := range args[1:] {
		value, _ := strconv.Atoi(arg)
		values = append(values, int32(value))
	}
	arg, _ := strconv.Atoi(args[0])
	testType := TestType(arg)

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
	client := NewBetterTestClient(conn)
	resp, err := client.Add(context.Background(), &BetterNumericRequest{Type: testType, Values: values})
	if err != nil {
		log.Fatalln("failed server call:", err)
	}
	log.Println("resp:", resp)
}
