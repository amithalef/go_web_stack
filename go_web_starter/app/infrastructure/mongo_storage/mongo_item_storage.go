package mongo_storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/amithnair91/go_web_stack/go_web_starter/app/domain"
)

type MongoItemStorage struct {
	collection *mongo.Collection
}

func (s MongoItemStorage) Save(item *domain.Item) (domain.Item, error) {
	_, err := s.collection.InsertOne(defaultContext(), item)
	if err != nil {
		return domain.Item{}, errors.New(fmt.Sprintf("Could not save to database: %s", err.Error()))
	}
	return *item, nil
}

func (s MongoItemStorage) Exists(id uuid.UUID) (bool, error) {
	filter := bson.M{"id": id}
	var result domain.Item
	if err := s.collection.FindOne(defaultContext(), filter).Decode(&result); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return false, nil
		}
		return false, errors.New(fmt.Sprintf("error while decoding item %s :", err.Error()))
	}
	return true, nil
}

//TODO figure out if  this needs to be instantiated every single time
func defaultContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx
}

func NewMongoItemStorage(database *mongo.Database) MongoItemStorage {
	collection := database.Collection("item")
	return MongoItemStorage{collection: collection}
}

/*
func (s MongoItemStorage) Upsert(item *domain.Item) (domain.Item, error) {
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"id", item.Id}}
	update := bson.D{{"$set", bson.D{{"name", item.Name}, {}}}}
	result, err := s.collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		return domain.Item{}, errors.New(fmt.Sprintf("Could not save to database: %s", err.Error()))
	}
	if result.UpsertedCount == 1 {
		return *item, nil
	}
	return domain.Item{}, errors.New(fmt.Sprintf("failed to Save to database %s", item.Id))
}
*/
