package config_test

import (
	"github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/testbuilders"
	"github.com/stretchr/testify/assert"
	"testing"
)

var configBuilder = testbuilders.ConfigBuilder{}

func TestValidateConfigWithCorrectConfig(t *testing.T) {
	config := configBuilder.Build()
	err := config.Validate()

	assert.Nil(t,err)
}

func TestValidateReturnsErrorIfMongoHostIsNotSet(t *testing.T) {
	config := configBuilder.Build()
	config.MONGO_HOST = ""

	err := config.Validate()

	assert.NotNil(t, err)
	assert.Equal(t, "environment variable MONGO_HOST is not set", err.Error())
}

func TestValidateReturnsErrorIfMongoPortIsNotSet(t *testing.T) {
	config := configBuilder.Build()
	config.MONGO_PORT = ""
	err := config.Validate()

	assert.NotNil(t, err)
	assert.Equal(t, "environment variable MONGO_PORT is not set", err.Error())
}

func TestValidateReturnsErrorIfMongoDatabaseIsNotSet(t *testing.T) {
	config := configBuilder.Build()
	config.MONGO_DATABASE = ""

	err := config.Validate()

	assert.NotNil(t, err)
	assert.Equal(t, "environment variable MONGO_DATABASE is not set", err.Error())
}

func TestValidateReturnsErrorIfMongoIfPortIsNotSet(t *testing.T) {
	config := configBuilder.Build()
	config.APP_PORT = ""

	err := config.Validate()

	assert.NotNil(t, err)
	assert.Equal(t, "environment variable APP_PORT is not set", err.Error())
}
