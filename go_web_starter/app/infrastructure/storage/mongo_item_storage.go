package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/amithnair91/go_web_stack/go_web_starter/app/domain"
)

type MongoItemStorage struct {
	collection *mongo.Collection
	context    context.Context
}

func (s MongoItemStorage) Save(item *domain.Item) (string, error) {
	one, err := s.collection.InsertOne(s.context, item)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Could not save to database: %s", err.Error()))
	}
	return one.InsertedID.(primitive.ObjectID).String(), nil
}

func (s MongoItemStorage) Exists(id uuid.UUID) (bool, error) {
	filter := bson.M{"id": id}
	var result domain.Item
	if err := s.collection.FindOne(s.context, filter).Decode(&result); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return false, nil
		}
		return false, errors.New(fmt.Sprintf("error while decoding item %s :", err.Error()))
	}
	return true, nil
}

func NewMongoItemStorage(database *mongo.Database) MongoItemStorage {
	collection := database.Collection("item")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return MongoItemStorage{collection: collection, context: ctx}
}
