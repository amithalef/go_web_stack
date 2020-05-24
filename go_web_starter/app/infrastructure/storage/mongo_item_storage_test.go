package storage_test

import (
	"github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/test_utils"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"

	"github.com/amithnair91/go_web_stack/go_web_starter/app/domain"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/storage"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) (test_utils.MongoTestContainer, *mongo.Database) {
	return test_utils.StartMongoDbForTest(t)
}

func teardown(container test_utils.MongoTestContainer) {
	container.Stop()
}

func TestSaveItemToDatabase(t *testing.T) {
	container, database := setup(t)
	itemStorage := storage.NewMongoItemStorage(database)
	bag, _ := domain.NewItem("bag")

	id, err := itemStorage.Save(bag)

	assert.Nil(t, err)
	assert.NotNil(t, id)
	teardown(container)
}

func TestExistsReturnsFalseIfNotExistsInDatabase(t *testing.T) {
	container, database := setup(t)
	itemStorage := storage.NewMongoItemStorage(database)
	bag, _ := domain.NewItem("bag")

	exists, err := itemStorage.Exists(bag.Id)

	assert.Nil(t, err)
	assert.False(t, exists)
	teardown(container)
}

func TestExistsReturnsTrueIfExistsInDatabase(t *testing.T) {
	container, database := setup(t)
	itemStorage := storage.NewMongoItemStorage(database)
	bag, _ := domain.NewItem("bag")
	itemStorage.Save(bag)

	exists, err := itemStorage.Exists(bag.Id)

	assert.Nil(t, err)
	assert.True(t, exists)
	teardown(container)
}
