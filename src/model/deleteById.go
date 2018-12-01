package model

import (
	c "app/src/constants"
	pb "app/src/grpc"
	"context"
)

func (rpc *RPC) DeleteById(context context.Context, request *pb.ByIdRequest) (*pb.Response, error) {

	return rpc.Update(context, &pb.UpdateRequest{
		Data: &pb.UpdateRequestData{
			Type: request.Data.Type,
			Id:   request.Data.Id,
			Attributes: map[string]string{
				"status": c.STATUS_DELETED,
			},
		},
	})
}
