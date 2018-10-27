package catalog

import (
	pb "app/grpc"
	"app/mongo"
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

func (rpc *RPC) Create(context context.Context, request *pb.CreateRequest) (*pb.Response, error) {
	log.Printf("REQUEST: %v", request)
	dataType := request.Data.Type
	dataAttributes := request.Data.Attributes
	collection, err := mongo.GetCollection(dataType)
	if err != nil {
		return nil, err
	}

	var instanceExist map[string]interface{}
	if err := collection.Find(dataAttributes).One(&instanceExist); err != nil {
		if err.Error() == "not found" {
			if err := collection.Insert(dataAttributes); err != nil {
				return nil, status.Errorf(codes.Internal, err.Error())
			}
		}
	}

	//if instanceExist != nil {
	//return nil, status.Errorf(codes.AlreadyExists, "Already exists")
	//}

	var newInstance map[string]interface{}
	if err := collection.Find(dataAttributes).One(&newInstance); err != nil {
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
