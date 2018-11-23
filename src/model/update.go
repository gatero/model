package model

import (
	pb "app/src/grpc"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

func (rpc *RPC) Update(context context.Context, request *pb.UpdateRequest) (*pb.Response, error) {
	if err := rpc.Mongo.SetCollection(request.Data.Type); err != nil {
		return nil, err
	}
	defer rpc.Mongo.Session.Close()

	attributes := request.Data.Attributes
	query := bson.M{
		"_id": bson.ObjectIdHex(request.Data.Id),
	}

	instanceToUpsert := make(bson.M)
	for key, _ := range attributes {
		if key != "" && key != "_id" {
			instanceToUpsert[key] = attributes[key]
		}
	}

	_, err := rpc.Mongo.Collection.Upsert(query, bson.M{"$set": instanceToUpsert})
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	updated, err := rpc.FindById(context, &pb.FindByIdRequest{
		Data: &pb.FindByIdRequestData{
			Type: request.Data.Type,
			Id:   request.Data.Id,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.Response{
		Type:       request.Data.Type,
		Id:         request.Data.Id,
		Attributes: updated.Attributes,
	}, nil
}
