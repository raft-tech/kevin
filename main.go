package main

import (
	"kevin/pingpong"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}
	grpcServer := grpc.NewServer()
	pingpong.RegisterPongServiceServer(grpcServer, &pingpong.Server{})
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over on port 9000: %v", err)
	}

}
