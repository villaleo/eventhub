// Package main implements a RouteGuide server.
package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	"github.com/joho/godotenv"
	pb "github.com/villaleo/eventhub/eventhub"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 50051, "The server port")

// Server wraps the gRPC server and its dependencies.
type Server struct {
	pb.UnimplementedEventManagerServer

	logger *zap.Logger
	db     *mongo.Client
}

func main() {
	flag.Parse()
	godotenv.Load()

	app := fx.New(
		fx.Provide(
			NewZapLogger,
			NewServer,
			grpc.NewServer,
		),
		fx.Invoke(
			Register,
			StartServer,
			ConnectDatabase,
		),
	)
	app.Run()
}

// NewZapLogger initializes a Zap logger.
func NewZapLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}

// NewServer creates a new instance of Server.
//
// By default, Server.db is nil.
func NewServer(logger *zap.Logger) *Server {
	return &Server{logger: logger}
}

// Register the gRPC server with the provided logger and listener.
func Register(grpcSrv *grpc.Server, server *Server) {
	pb.RegisterEventManagerServer(grpcSrv, server)
}

// StartServer is the lifecycle hook to start the gRPC server.
func StartServer(lc fx.Lifecycle, grpcSrv *grpc.Server, logger *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			addr := fmt.Sprintf("localhost:%d", *port)
			lis, err := net.Listen("tcp", addr)
			if err != nil {
				logger.Fatal("failed to listen", zap.Error(err))
			}

			logger.Info("grpc server started", zap.String("address", addr))
			go func() {
				if err := grpcSrv.Serve(lis); err != nil {
					logger.Fatal("failed to serve", zap.Error(err))
				}
			}()

			return nil
		},
		OnStop: func(_ context.Context) error {
			grpcSrv.GracefulStop()
			logger.Info("shutting down the server")
			return nil
		},
	})
}
