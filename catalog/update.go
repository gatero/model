package catalog

import (
	pb "app/grpc"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

func (rpc *RPC) Update(context context.Context, request *pb.UpdateRequest) (*pb.Response, error) {
	rpc.Mongo.CollectionName = request.Data.Type
	if err := rpc.Mongo.SetCollection(); err != nil {
		return nil, err
	}
	defer rpc.Mongo.Session.Close()

	dataAttributes := request.Data.Attributes
	query := bson.M{
		"_id": bson.ObjectIdHex(request.Data.Type),
	}

	var instanceToUpdate map[string]interface{}
	if err := rpc.Mongo.Collection.Find(query).One(&instanceToUpdate); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if instanceToUpdate == nil {
		return nil, status.Errorf(codes.NotFound, "Not found")
	}

	for key, _ := range dataAttributes {
		if key != "" && key != "_id" {
			instanceToUpdate[key] = dataAttributes[key]
		}
	}

	if err := rpc.Mongo.Collection.Update(query, &instanceToUpdate); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	outputAttributes := make(map[string]string)
	for key, value := range instanceToUpdate {
		if value != nil && key != "_id" {
			outputAttributes[key] = value.(string)
		}
	}

	return &pb.Response{
		Type:       request.Data.Type,
		Id:         request.Data.Id,
		Attributes: outputAttributes,
	}, nil
}
