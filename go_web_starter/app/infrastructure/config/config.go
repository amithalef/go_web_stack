package config

import "errors"

type Config struct {
	MONGO_URI string `env:"MONGO_URI"`
	APP_PORT  string `env:"APP_PORT" envDefault:"8888"`
}

func (c Config) Validate() error {
	if len(c.MONGO_URI) < 1 {
		return errors.New("environment variable MONGO_URI is not set")
	}
	if len(c.APP_PORT) < 1 {
		return errors.New("environment variable APP_PORT is not set")
	}
	return nil
}
