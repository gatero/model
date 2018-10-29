package catalog

import (
	pb "app/grpc"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

func (rpc *RPC) Create(context context.Context, request *pb.CreateRequest) (*pb.Response, error) {
	rpc.Mongo.CollectionName = request.Data.Type
	if err := rpc.Mongo.SetCollection(); err != nil {
		return nil, err
	}
	defer rpc.Mongo.Session.Close()

	Attributes := request.Data.Attributes

	instance := make(map[string]interface{})
	instance["_id"] = bson.NewObjectId()
	for key, value := range Attributes {
		if value != "" && key != "_id" {
			instance[key] = value
		}
	}

	if err := rpc.Mongo.Collection.Insert(instance); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.Response{
		Type:       request.Data.Type,
		Id:         instance["_id"].(bson.ObjectId).Hex(),
		Attributes: Attributes,
	}, nil
}
