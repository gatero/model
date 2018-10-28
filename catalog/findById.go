package catalog

import (
	pb "app/grpc"
	"app/mongo"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

func (rpc *RPC) FindById(context context.Context, request *pb.FindByIdRequest) (*pb.Response, error) {
	dataType := request.Data.Type
	dataId := request.Data.Id
	query := bson.M{
		"_id": bson.ObjectIdHex(dataId),
	}

	collection, err := mongo.GetCollection(dataType)
	if err != nil {
		return nil, err
	}

	var instancetoUpdate map[string]interface{}
	if err := collection.Find(query).One(&instancetoUpdate); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if instancetoUpdate == nil {
		return nil, status.Errorf(codes.NotFound, "Not found")
	}

	outputAttributes := make(map[string]string)
	for key, value := range instancetoUpdate {
		if value != nil && key != "_id" {
			outputAttributes[key] = value.(string)
		}
	}

	return &pb.Response{
		Type:       dataType,
		Id:         dataId,
		Attributes: outputAttributes,
	}, nil
}
