package catalog

import (
	pb "app/grpc"
	"context"
)

func (rpc *RPC) FindById(context context.Context, request *pb.FindByIdRequest) (*pb.Response, error) {
	return &pb.Response{}, nil
}
