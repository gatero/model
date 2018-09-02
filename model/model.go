package model

import (
	pb "app/grpc"
	"context"
	"log"
)

type RPC struct{}

func (rpc *RPC) Create(context context.Context, request *pb.CreateRequest) (*pb.Response, error) {
	log.Printf("CREATE REQUEST: %v", request)
	return &pb.Response{}, nil
}

func (rpc *RPC) Update(context context.Context, request *pb.UpdateRequest) (*pb.Response, error) {
	return &pb.Response{}, nil
}

func (rpc *RPC) FindOne(context context.Context, request *pb.FindOneRequest) (*pb.Response, error) {
	return &pb.Response{}, nil
}

func (rpc *RPC) Find(context context.Context, request *pb.FindRequest) (*pb.FindResponse, error) {
	return &pb.FindResponse{}, nil
}

func (rpc *RPC) DeleteById(context context.Context, request *pb.DeleteByIdRequest) (*pb.Response, error) {
	return &pb.Response{}, nil
}
