package catalog

import (
	pb "app/grpc"
	"app/mongo"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RPC struct{}

func (rpc *RPC) Create(context context.Context, request *pb.CreateRequest) (*pb.Response, error) {
	dataType := request.Data.Type
	instance := request.Data.Attributes
	collection, err := mongo.GetCollection(request.Data.Type)
	if err != nil {
		return nil, err
	}

	if err := collection.Insert(instance); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

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
