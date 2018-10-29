package main

import (
	"log"
	"net"

	"app/catalog"
	pb "app/grpc"

	c "app/constants"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":"+c.CATALOG_PORT)
	if err != nil {
		log.Fatalf("ERROR: Failed listening %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterCatalogServer(server, &catalog.RPC{})

	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("ERROR: Failed to serve %v", err)
	}
}
