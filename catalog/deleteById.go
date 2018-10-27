package catalog

import (
	pb "app/grpc"
	"context"
)

func (rpc *RPC) DeleteById(context context.Context, request *pb.DeleteByIdRequest) (*pb.Response, error) {
	return &pb.Response{}, nil
}
