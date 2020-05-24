package storage_test

import (
	"github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/test_utils"
	"testing"

	"github.com/amithnair91/go_web_stack/go_web_starter/app/domain"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/storage"
	"github.com/stretchr/testify/assert"
)

func TestSaveItemToDatabase(t *testing.T) {
	container := test_utils.MongoTestContainer{}
	defer container.Stop()
	container.Start(t)
	database, _ := storage.Connect(container.IP, container.Port, test_utils.DatabaseName)
	itemStorage := storage.NewMongoItemStorage(database)
	bag, _ := domain.NewItem("bag")

	id, err := itemStorage.Save(bag)

	assert.Nil(t, err)
	assert.NotNil(t, id)
}

func TestExistsReturnsFalseIfNotExistsInDatabase(t *testing.T) {
	container := test_utils.MongoTestContainer{}
	defer container.Stop()
	container.Start(t)
	database, _ := storage.Connect(container.IP, container.Port, test_utils.DatabaseName)
	itemStorage := storage.NewMongoItemStorage(database)
	bag, _ := domain.NewItem("bag")

	exists, err := itemStorage.Exists(bag.Id)

	assert.Nil(t, err)
	assert.False(t, exists)
}

func TestExistsReturnsTrueIfExistsInDatabase(t *testing.T) {
	container := test_utils.MongoTestContainer{}
	defer container.Stop()
	container.Start(t)
	database, _ := storage.Connect(container.IP, container.Port, test_utils.DatabaseName)
	itemStorage := storage.NewMongoItemStorage(database)
	bag, _ := domain.NewItem("bag")
	itemStorage.Save(bag)

	exists, err := itemStorage.Exists(bag.Id)

	assert.Nil(t, err)
	assert.True(t, exists)
}
