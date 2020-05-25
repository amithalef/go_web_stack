package testbuilders

import "github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/config"

type ConfigBuilder struct {
	mongoPort       string
	mongoHost       string
	mongoDatabase   string
	applicationPort string
}

func (cb *ConfigBuilder) WithMongoPort(port string) *ConfigBuilder {
	cb.mongoPort = port
	return cb
}

func (cb *ConfigBuilder) WithMongoHost(host string) *ConfigBuilder {
	cb.mongoHost = host
	return cb
}

func (cb *ConfigBuilder) WithMongodatabase(database string) *ConfigBuilder {
	cb.mongoDatabase = database
	return cb
}

func (cb *ConfigBuilder) WithapplicationPort(applicationPort string) *ConfigBuilder {
	cb.applicationPort = applicationPort
	return cb
}

func (cb *ConfigBuilder) Build() config.Config {
	if cb.mongoPort == "" {
		cb.mongoPort = "27017"
	}
	if cb.mongoDatabase == "" {
		cb.mongoDatabase = "test"
	}
	if cb.mongoHost == "" {
		cb.mongoHost = "localhost"
	}
	if cb.applicationPort == "" {
		cb.applicationPort = "8888"
	}
	return config.Config{
		MONGO_HOST:     cb.mongoHost,
		MONGO_PORT:     cb.mongoHost,
		MONGO_DATABASE: cb.mongoDatabase,
		APP_PORT:       cb.applicationPort,
	}
}
