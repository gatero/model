package catalog

import (
	pb "app/grpc"
	"context"
)

func (rpc *RPC) FindOne(context context.Context, request *pb.FindOneRequest) (*pb.Response, error) {
	return &pb.Response{}, nil
}
