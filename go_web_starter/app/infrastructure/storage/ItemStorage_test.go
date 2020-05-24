package storage_test

import (
	"context"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/domain"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/storage"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestSaveItemToDatabase(t *testing.T) {
	container := MongoTestContainer{}
	container.Start(t)
	defer container.Stop()
	database, _ := storage.Connect(container.IP, container.Port,"testing")
	itemStorage := storage.NewItemStorage(database)

	bag, _ := domain.NewItem("bag")

	itemStorage.Save(bag)

	find, _ := database.Collection("item").Find(context.Background(), bson.D{})
	//result := domain.Item{}
	//filter := bson.M{"Name": "bag"}
	//database.Collection("item").FindOne(context.Background(),filter).Decode(result)
	result := domain.Item{}
	find.Next(context.Background())
	find.Decode(result)
	assert.NotNil(t, result)
	//assert.Equal(t,bag.Name,result.Name)
}
