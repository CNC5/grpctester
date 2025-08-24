package main

import (
    "fmt"
    "flag"
  	"context"
  	"log"
  	"net"

  	//pb "grpctester/helloworld" // generated from proto
  	pb "grpctester/gen" // generated from proto
  	"google.golang.org/grpc"
)

type server struct {
  	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	  return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
    port := flag.String("port", "50051", "Port to listen on")
    flag.Parse()

    address := fmt.Sprintf(":%s", *port)
  	lis, err := net.Listen("tcp", address)
  	if err != nil {
    		log.Fatalf("failed to listen: %v", err)
  	}

  	grpcServer := grpc.NewServer()
  	pb.RegisterGreeterServer(grpcServer, &server{})
  	log.Printf("gRPC server listening on %s\n", address)
  	if err := grpcServer.Serve(lis); err != nil {
    		log.Fatalf("failed to serve: %v", err)
  	}
}
