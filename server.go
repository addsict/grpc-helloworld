package main

import (
	"fmt"
	pb "localhost/helloworld/proto"
	"log"
	"net"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Greet(ctx context.Context, in *pb.RequestType) (*pb.ResponseType, error) {
	return &pb.ResponseType{Msg: fmt.Sprintf("Hello GRPC and %s!", in.Name)}, nil
}

func main() {
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s.Serve(lis)
}
