package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ItemStorage struct {
	collection *mongo.Collection
	context    context.Context
}

func (s ItemStorage) Save(item *domain.Item) (string, error) {
	one, err := s.collection.InsertOne(s.context, item)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Could not save to database: %s", err.Error()))
	}
	return one.InsertedID.(primitive.ObjectID).String(), nil
}

func NewItemStorage(database *mongo.Database) ItemStorage {
	collection := database.Collection("item")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return ItemStorage{collection: collection, context: ctx}
}
