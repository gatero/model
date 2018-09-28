
import (
	pb "app/grpc"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (rpc *RPC) Update(context context.Context, request *pb.UpdateRequest) (*pb.Response, error) {
	dataType := request.Data.Type
	dataId := request.Data.Id
	dataAttributes := request.Data.Attributes
	collection, err := mongo.GetCollection(dataType)
	if err != nil {
		return nil, err
	}

	var instancetoUpdate map[string]interface{}
	if err := collection.Find(dataAttributes).One(&instancetoUpdate); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if instancetoUpdate == nil {
		return nil, status.Errorf(codes.NotFound, "Not found")
	}

	for key, value := range dataAttributes {
		if key != nil && key != "_id" {
			instancetoUpdate[key] = dataAttributes[key].(string)
		}
	}

	if err := collection.Update(query, &instancetoUpdate); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	outputAttributes := instancetoUpdate

	return &pb.Response{
		Type:       dataType,
		Id:         dataId,
		Attributes: outputAttributes,
	}, nil
}

