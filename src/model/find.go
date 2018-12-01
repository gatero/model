package model

import (
	pb "app/src/grpc"
	"context"
	"encoding/json"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

func (rpc *RPC) Find(context context.Context, request *pb.FindRequest) (*pb.FindResponse, error) {
	if err := rpc.Mongo.SetCollection(request.Type); err != nil {
		return nil, err
	}
	defer rpc.Mongo.Session.Close()

	page := int32(1)
	if request.Page > 1 {
		page = request.Page
	}

	limit := int32(10)
	if request.Limit != 0 {
		limit = request.Limit
	}

	sortBy := "ASC"
	if request.SortBy != "" {
		sortBy = request.SortBy
	}

	if request.Type == "" {
		return nil, status.Errorf(codes.Internal, "Type can not be empty")
	}

	filter := make(bson.M)
	if request.Filter != nil {
		if err := json.Unmarshal([]byte(request.Filter), &filter); err != nil {
			return nil, err
		}

		// this range makes generic search's
		for key, value := range filter {
			filter[key] = bson.RegEx{value.(string), ""}
		}
	}

	offset := int((page - 1) * limit)

	var instances []bson.M

	if err := rpc.Mongo.Collection.Find(filter).Sort(sortBy).Skip(offset).Limit(int(limit)).All(&instances); err != nil {
		return nil, err
	}

	var data []*pb.Response

	for _, instance := range instances {
		formated := &pb.Response{}
		formated.Type = request.Type
		formated.Id = instance["_id"].(bson.ObjectId).Hex()
		delete(instance, "_id")

		formated.Attributes = (func() map[string]string {
			attributes := make(map[string]string)

			for key, value := range instance {
				attributes[key] = value.(string)
			}

			return attributes
		})()

		data = append(data, formated)
	}

	return &pb.FindResponse{
		Metadata: &pb.Metadata{
			Paginate: &pb.Paginate{
				Page:   int32(page),
				Limit:  int32(limit),
				SortBy: sortBy,
			},
		},
		Data: data,
	}, nil
}
