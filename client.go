package main

import (
    "flag"
    "fmt"
  	"context"
  	"log"
  	"time"

  	pb "grpctester/gen"
  	"google.golang.org/grpc"
)

func main() {
    port := flag.String("port", "50051", "Port to listen on")
    flag.Parse()

  	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", *port), grpc.WithInsecure())
  	if err != nil {
    		log.Fatalf("did not connect: %v", err)
  	}
  	defer conn.Close()

  	c := pb.NewGreeterClient(conn)

  	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
  	defer cancel()

  	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "World"})
  	if err != nil {
    		log.Fatalf("could not greet: %v", err)
  	}
  	log.Printf("Greeting: %s", r.Message)
}
