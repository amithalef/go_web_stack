package storage_test

import (
	"github.com/amithnair91/go_web_stack/go_web_starter/app/domain"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveItemToDatabase(t *testing.T) {
	container := MongoTestContainer{}
	defer container.Stop()
	container.Start(t)
	database, _ := storage.Connect(container.IP, container.Port, "testing")
	itemStorage := storage.NewItemStorage(database)
	bag, _ := domain.NewItem("bag")

	id, err := itemStorage.Save(bag)

	assert.NotNil(t, id)
	assert.Nil(t, err)
}
