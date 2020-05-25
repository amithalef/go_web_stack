package config_test

import (
	"github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateConfigWithCorrectConfig(t *testing.T) {
	config := config.Config{MONGO_URI: "MONGO_URI", APP_PORT: "3333"}
	config.Validate()
}

func TestValidateReturnsErrorIfMongoURIIsNotSet(t *testing.T) {
	config := config.Config{MONGO_URI: "", APP_PORT: "3333"}
	err := config.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "environment variable MONGO_URI is not set", err.Error())
}

func TestValidateReturnsDefaultPortIfPortIsNotSet(t *testing.T) {
	config := config.Config{MONGO_URI: "mongodb://url:port", APP_PORT: ""}

	err := config.Validate()

	assert.NotNil(t, err)
	assert.Equal(t, "environment variable APP_PORT is not set", err.Error())
}
