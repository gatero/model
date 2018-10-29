package catalog

import (
	pb "app/grpc"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

func (rpc *RPC) FindById(context context.Context, request *pb.FindByIdRequest) (*pb.Response, error) {
	rpc.Mongo.CollectionName = request.Data.Type
	if err := rpc.Mongo.SetCollection(); err != nil {
		return nil, err
	}
	defer rpc.Mongo.Session.Close()

	query := bson.M{
		"_id": bson.ObjectIdHex(request.Data.Id),
	}

	var result map[string]interface{}
	if err := rpc.Mongo.Collection.Find(query).One(&result); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if result == nil {
		return nil, status.Errorf(codes.NotFound, "Not found")
	}

	outputAttributes := make(map[string]string)
	for key, value := range result {
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
