package storage_test

import (
	"github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnectMongo(t *testing.T) {
	container := MongoTestContainer{}
	container.Start(t)
	defer container.Stop()
	database, err := storage.Connect(container.IP, container.Port, "testing")
	assert.Nil(t, err)
	assert.NotNil(t, database)
	collection := database.Collection("numbers")
	assert.Equal(t, "numbers", collection.Name())
}

func TestConnectMongoReturnsErrorWhenIpIsEmpty(t *testing.T) {
	container := MongoTestContainer{}
	container.Start(t)
	defer container.Stop()
	_, err := storage.Connect("", container.Port, "testing")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "ip cannot be empty")
}

func TestConnectMongoReturnsErrorWhenPortIsEmpty(t *testing.T) {
	container := MongoTestContainer{}
	container.Start(t)
	defer container.Stop()
	_, err := storage.Connect(container.IP, "", "testing")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "port cannot be empty")
}

func TestConnectMongoReturnsErrorWhenDatabaseNameIsEmpty(t *testing.T) {
	container := MongoTestContainer{}
	container.Start(t)
	defer container.Stop()
	_, err := storage.Connect(container.IP, container.Port, "")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "databaseName cannot be empty")
}

func TestShouldReturnWhenFailsToConnectToMongo(t *testing.T) {
	_, err := storage.Connect("non-existant-ip", "8080", "testing")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "could not ping mongo db :")
}
