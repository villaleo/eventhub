package main

import (
	"context"
	"reflect"
	"strings"

	pb "github.com/villaleo/eventhub/eventhub"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	errInternal = status.Error(codes.Internal, "an internal error occurred")
)

func (s *Server) NewEvent(ctx context.Context, event *pb.Event) (*pb.Event, error) {
	c := s.db.Database(dbEvents).Collection(dbcEventHub)

	// Use reflection to create a bson.D excluding the `Id` field
	doc := bson.D{}
	val := reflect.ValueOf(event).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		// Skip the field if it's unexported or named `Id`
		if !field.IsExported() || field.Name == "Id" {
			continue
		}
		// Rename field to begin with a lowercase letter
		key := strings.ToLower(string(field.Name[0])) + field.Name[1:]
		doc = append(doc, bson.E{Key: key, Value: val.Field(i).Interface()})
	}

	// Insert the document into the collection
	res, err := c.InsertOne(ctx, doc)
	if err != nil {
		s.logger.Error("failed to insert new event", zap.Any("event", event), zap.Error(err))
		return nil, errInternal
	}

	// Set the Id of the Event to be returned
	switch id := res.InsertedID.(type) {
	case primitive.ObjectID:
		event.Id = id.Hex()
		s.logger.Info("inserted new event", zap.Any("event", event), zap.Any("id", id.String()))
	default:
		s.logger.Error("unexpected generated id type", zap.Any("id", id))
		return nil, errInternal
	}

	return event, nil
}

func (s *Server) UpdateEvent(ctx context.Context, event *pb.Event) (*pb.Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}

func (s *Server) DeleteEvent(ctx context.Context, event *pb.Event) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}

func (s *Server) ListEvents(req *pb.ListEventsRequest, stream pb.EventManager_ListEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListEvents not implemented")
}

func (s *Server) FindEvents(req *pb.FindEventsRequest, stream pb.EventManager_FindEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method FindEvents not implemented")
}
