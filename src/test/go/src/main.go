package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "contract"
	"google.golang.org/grpc"
)

const (
	address     = "172.25.7.73:50051"
	defaultName = "world"
)

func main() {

	log.Printf("Started")

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	log.Printf("Connected")

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.Printf("Say Hello")
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Get Hello response")

	log.Printf("Greeting: %s", r.GetMessage())
}
