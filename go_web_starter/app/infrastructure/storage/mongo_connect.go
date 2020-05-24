package storage

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func Connect(ip string, port string, databaseName string) (*mongo.Database, error) {
	if len(ip) < 1 {
		return nil, errors.New("ip cannot be empty")
	}
	if len(port) < 1 {
		return nil, errors.New("port cannot be empty")
	}
	if len(databaseName) < 1 {
		return nil, errors.New("databaseName cannot be empty")
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s/%s", ip, port, databaseName)))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not create mongo client : %s", err))
	}
	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not ping mongo db : %s", err))
	}
	return client.Database(databaseName), nil
}
