package main

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ConnectDatabase creates a ready to use client connection to a MongoDB
// cluster and stores it in srv.db
func ConnectDatabase(lc fx.Lifecycle, srv *Server, logger *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			uri := fmt.Sprintf(
				"mongodb+srv://%s:%s@%s",
				os.Getenv("MONGO_USERNAME"),
				os.Getenv("MONGO_PW"),
				os.Getenv("MONGO_URL"),
			)

			client, err := mongo.Connect(
				ctx,
				options.Client().
					ApplyURI(uri).
					SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1)),
			)
			if err != nil {
				logger.Fatal("failed to connect to mongo cluster", zap.Error(err))
			}

			if err = client.Ping(ctx, nil); err != nil {
				logger.Fatal("failed to ping mongo cluster", zap.Error(err))
			}
			srv.db = client

			logger.Info("connected to mongo cluster")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := srv.db.Disconnect(ctx); err != nil {
				logger.Fatal("failed to disconnect from cluster", zap.Error(err))
			}

			logger.Info("disconnecting from mongo cluster")
			return nil
		},
	})
}
