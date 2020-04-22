package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc_test/proto"
	"log"
	"os"
)

const (
	address     = "localhost:50052"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewToUpperClient(conn)

	// Contact the server and print out its response.
	name := "hello world"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r1, err1 := c.Greeting(context.Background(), &pb.UpperRequest{Name: "hello"})
	if err1 != nil {
		log.Fatalf("could not gretting", err)
	} else {
		log.Printf("response: %s", r1.Message)
	}

	r, err := c.Upper(context.Background(), &pb.UpperRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	data, err := r.Recv()
	log.Printf("Response: %s", data.Data)
}
