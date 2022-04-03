package database

import (
	"context"
	"dysn/character/internal/config"
	"dysn/character/internal/service/logger"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongodb struct {
	lgr      *logger.Logger
	client   *mongo.Client
	Database *mongo.Database
}

const dbUri = "mongodb://%s:%s/"

func Init(ctx context.Context, cfg *config.Config, lgr *logger.Logger) *Mongodb {
	uri := fmt.Sprintf(dbUri, cfg.GetHost(), cfg.GetPort())

	clientOptions := options.Client().ApplyURI(uri).
		SetAuth(options.Credential{
			AuthSource: cfg.GetDbName(),
			Username:   cfg.GetUser(),
			Password:   cfg.GetPassword(),
		})
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		lgr.ErrorLog.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		lgr.ErrorLog.Fatal(err)
	}

	lgr.InfoLog.Println("MongoDb has started")

	return &Mongodb{
		client:   client,
		Database: client.Database(cfg.GetDbName()),
	}
}

func (m *Mongodb) GetClient() *mongo.Client {
	return m.client
}
