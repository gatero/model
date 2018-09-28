package catalog

import (
	pb "app/grpc"
	"app/mongo"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

type RPC struct{}

func (rpc *RPC) Create(context context.Context, request *pb.CreateRequest) (*pb.Response, error) {
	dataType := request.Data.Type
	instance := request.Data.Attributes
	collection, err := mongo.GetCollection(dataType)
	if err != nil {
		return nil, err
	}

	var instanceExist map[string]interface{}
	if err := collection.Find(instance).One(&instanceExist); err != nil {
		if err.Error() == "not found" {
			if err := collection.Insert(instance); err != nil {
				return nil, status.Errorf(codes.Internal, err.Error())
			}
		}
	}

	if instanceExist != nil {
		return nil, status.Errorf(codes.AlreadyExists, "Already exists")
	}

	var newInstance map[string]interface{}
	if err := collection.Find(instance).One(&newInstance); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	outputAttributes := make(map[string]string)
	for key, value := range newInstance {
		if value != nil && key != "_id" {
			outputAttributes[key] = value.(string)
		}
	}

	return &pb.Response{
		Type:       dataType,
		Id:         newInstance["_id"].(bson.ObjectId).Hex(),
		Attributes: outputAttributes,
	}, nil
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
