package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "grpc_test/proto"
	"log"
	"net"
)

const (
	port = ":50052"
)

type server struct{
}

func (s *server) Upper(req *pb.UpperRequest, srv pb.ToUpper_UpperServer) error {
	log.Printf("Received: %s", req.Name)
	srv.Send(&pb.Data{Data:"fuck " + req.Name+"!"})
	return nil
}

func (s *server) Greeting(ctx context.Context, in *pb.UpperRequest) (*pb.UpperReply, error) {
	log.Printf("Received: %s", in.Name)
	return &pb.UpperReply{Message: "hello i am kiko!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterToUpperServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}