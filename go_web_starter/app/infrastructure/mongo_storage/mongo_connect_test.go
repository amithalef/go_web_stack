package mongo_storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/mongo_storage"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/test_utils"
)

func TestConnectMongo(t *testing.T) {
	container := test_utils.MongoTestContainer{}
	container.Start()
	defer container.Stop()
	database, err := mongo_storage.Connect(container.IP, container.Port, "testing")
	assert.Nil(t, err)
	assert.NotNil(t, database)
	collection := database.Collection("numbers")
	assert.Equal(t, "numbers", collection.Name())
}

func TestConnectMongoReturnsErrorWhenIpIsEmpty(t *testing.T) {
	container := test_utils.MongoTestContainer{}
	container.Start()
	defer container.Stop()
	_, err := mongo_storage.Connect("", container.Port, "testing")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "ip cannot be empty")
}

func TestConnectMongoReturnsErrorWhenPortIsEmpty(t *testing.T) {
	container := test_utils.MongoTestContainer{}
	container.Start()
	defer container.Stop()
	_, err := mongo_storage.Connect(container.IP, "", "testing")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "port cannot be empty")
}

func TestConnectMongoReturnsErrorWhenDatabaseNameIsEmpty(t *testing.T) {
	container := test_utils.MongoTestContainer{}
	container.Start()
	defer container.Stop()
	_, err := mongo_storage.Connect(container.IP, container.Port, "")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "databaseName cannot be empty")
}

func TestShouldReturnWhenFailsToConnectToMongo(t *testing.T) {
	_, err := mongo_storage.Connect("non-existant-ip", "8080", "testing")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "could not ping mongo db :")
}
