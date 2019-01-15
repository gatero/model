package model

import (
	pb "app/src/grpc"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

func (rpc *RPC) Create(context context.Context, request *pb.CreateRequest) (*pb.Response, error) {
	if err := rpc.Mongo.SetCollection(request.Data.Type); err != nil {
		return nil, err
	}
	defer rpc.Mongo.Session.Close()

	options := request.Options
	attributes := request.Data.Attributes

	log.Println("hola mundo")
	if options != nil && options["unique"] == "true" {
		filter := fmt.Sprintf("{ \"name\": \"%s\" }", attributes["name"])
		results, err := rpc.Find(context, &pb.FindRequest{
			Type:   "profile",
			Filter: []byte(filter),
		})
		if err != nil {
			return nil, err
		}
		log.Printf("RESULTS: ", results)
	}

	instance := make(bson.M)
	instance["_id"] = bson.NewObjectId()
	for key, value := range attributes {
		if value != "" && key != "_id" {
			instance[key] = value
		}
	}

	if err := rpc.Mongo.Collection.Insert(instance); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.Response{
		Type:       request.Data.Type,
		Id:         instance["_id"].(bson.ObjectId).Hex(),
		Attributes: attributes,
	}, nil
}
