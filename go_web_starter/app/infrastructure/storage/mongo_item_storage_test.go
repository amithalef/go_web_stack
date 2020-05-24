package storage_test

import (
	"github.com/amithnair91/go_web_stack/go_web_starter/app/domain"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

const databaseName = "testing"

func TestSaveItemToDatabase(t *testing.T) {
	container := MongoTestContainer{}
	defer container.Stop()
	container.Start(t)
	database, _ := storage.Connect(container.IP, container.Port, databaseName)
	itemStorage := storage.NewMongoItemStorage(database)
	bag, _ := domain.NewItem("bag")

	id, err := itemStorage.Save(bag)

	assert.NotNil(t, id)
	assert.Nil(t, err)
}

func TestExistsReturnsFalseIfNotExistsInDatabase(t *testing.T) {
	container := MongoTestContainer{}
	defer container.Stop()
	container.Start(t)
	database, _ := storage.Connect(container.IP, container.Port, databaseName)
	itemStorage := storage.NewMongoItemStorage(database)
	bag, _ := domain.NewItem("bag")

	exists := itemStorage.Exists(bag.Id)

	assert.False(t, exists)
}
