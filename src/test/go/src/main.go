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
	defaultName = "world"
)

func main() {

	if len(os.Args) <= 3 {
		log.Printf("Not valid command, usage: /app 127.0.0.1 50051 hi")
		return
	}

	log.Printf("Started")

	host := os.Args[1]
	port := os.Args[2]
	name := os.Args[3]

	log.Printf("host %s", host)
	log.Printf("port %s", port)
	log.Printf("name %s", name)

	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	log.Printf("Connected")

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
