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
	client, err := storage.Connect(container.IP, container.Port)
	assert.NotNil(t, err)
	assert.NotNil(t, client)
	collection := client.Database("testing").Collection("numbers")
	assert.Equal(t, "numbers", collection.Name())
}
