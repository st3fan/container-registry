package main

import (
	"context"
	"log"
	"net"

	pb "github.com/st3fan/greeter-service/proto"
	"google.golang.org/grpc"
)

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *greeterServer) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{Message: "Hello " + req.Name}, nil
}

func newGreeterServer() *greeterServer {
	return &greeterServer{}
}

func main() {
	log.Println("Starting greeter-service on :9090")

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	pb.RegisterGreeterServer(server, newGreeterServer())
	if server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
