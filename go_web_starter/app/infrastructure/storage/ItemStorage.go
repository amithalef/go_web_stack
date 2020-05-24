package storage

import (
	"context"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ItemStorage struct {
	collection *mongo.Collection
	context    context.Context
}
// TODO return error on failure
func (s ItemStorage) Save(item *domain.Item) {
	s.collection.InsertOne(s.context, item)
}

func NewItemStorage(database *mongo.Database) ItemStorage {
	collection := database.Collection("item")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return ItemStorage{collection: collection, context: ctx}
}

