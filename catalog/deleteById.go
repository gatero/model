package catalog

import (
	pb "app/grpc"
	"context"
)

func (rpc *RPC) DeleteById(context context.Context, request *pb.DeleteByIdRequest) (*pb.Response, error) {

	return rpc.Update(context, &pb.UpdateRequest{
		Data: &pb.UpdateRequestData{
			Type: request.Data.Type,
			Id:   request.Data.Id,
			Attributes: map[string]string{
				"status": STATUS_DELETED,
			},
		},
	})
}
