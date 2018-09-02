package main

import (
	"log"
	"net"

	pb "app/grpc"
	"app/model"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":"+MODEL_PORT)
	if err != nil {
		log.Fatalf("ERROR: Failed listening %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterModelServer(server, &model.RPC{})

	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("ERROR: Failed to serve %v", err)
	}
}
