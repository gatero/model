package catalog

import (
	pb "app/grpc"
	"context"
)

type RPC struct{}

func (rpc *RPC) FindOne(context context.Context, request *pb.FindOneRequest) (*pb.Response, error) {
	return &pb.Response{}, nil
}

func (rpc *RPC) Find(context context.Context, request *pb.FindRequest) (*pb.FindResponse, error) {
	return &pb.FindResponse{}, nil
}

func (rpc *RPC) DeleteById(context context.Context, request *pb.DeleteByIdRequest) (*pb.Response, error) {
	return &pb.Response{}, nil
}
