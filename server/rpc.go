package main

import (
	"context"

	pb "github.com/villaleo/eventhub/eventhub"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) NewEvent(ctx context.Context, event *pb.Event) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
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
