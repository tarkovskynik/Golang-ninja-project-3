package mongo

import (
	"context"
	"fmt"

	"github.com/tarkovskynik/Golang-ninja-project-3/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB(ctx context.Context, cfg *config.Config) (*mongo.Client, error) {
	clientOptions := options.Client()
	clientOptions.SetAuth(options.Credential{
		Username: cfg.MongoDB.Username,
		Password: cfg.MongoDB.Password,
	})
	clientOptions.ApplyURI(cfg.MongoDB.URI)

	dbClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo: %w", err)
	}

	if err = dbClient.Ping(context.Background(), nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongo: %w", err)
	}

	return dbClient, nil
}
