package main

import (
	pb "localhost/helloworld/proto"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect server: %v", err)
	}
	defer con.Close()

	c := pb.NewHelloServiceClient(con)
	greet, err := c.Greet(context.Background(), &pb.RequestType{Name: "test"})
	if err != nil {
		log.Fatalf("failed to request: %v", err)
	}
	log.Printf("%s", greet.Msg)
}
