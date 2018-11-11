package crud

import (
	pb "app/src/grpc"
	"context"
)

func (rpc *RPC) Find(context context.Context, request *pb.FindRequest) (*pb.FindResponse, error) {
	return &pb.FindResponse{}, nil
}
