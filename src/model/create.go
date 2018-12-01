package model

import (
	pb "app/src/grpc"
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

func (rpc *RPC) Create(context context.Context, request *pb.CreateRequest) (*pb.Response, error) {
	if err := rpc.Mongo.SetCollection(request.Data.Type); err != nil {
		return nil, err
	}
	defer rpc.Mongo.Session.Close()
	headers, _ := metadata.FromIncomingContext(context)
	log.Printf("HEADERS: %v", headers)

	attributes := request.Data.Attributes

	instance := make(bson.M)
	instance["_id"] = bson.NewObjectId()
	for key, value := range attributes {
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
		Attributes: attributes,
	}, nil
}
