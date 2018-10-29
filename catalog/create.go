package catalog

import (
	pb "app/grpc"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

func (rpc *RPC) Create(context context.Context, request *pb.CreateRequest) (*pb.Response, error) {
	dataType := request.Data.Type
	dataAttributes := request.Data.Attributes

	rpc.Mongo.CollectionName = dataType
	if err := rpc.Mongo.SetCollection(); err != nil {
		return nil, err
	}
	defer rpc.Mongo.Session.Close()

	var instanceExist map[string]interface{}
	if err := rpc.Mongo.Collection.Find(dataAttributes).One(&instanceExist); err != nil {
		if err.Error() == "not found" {
			if err := rpc.Mongo.Collection.Insert(dataAttributes); err != nil {
				return nil, status.Errorf(codes.Internal, err.Error())
			}
		}
	}

	//if instanceExist != nil {
	//return nil, status.Errorf(codes.AlreadyExists, "Already exists")
	//}

	var newInstance map[string]interface{}
	if err := rpc.Mongo.Collection.Find(dataAttributes).One(&newInstance); err != nil {
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
