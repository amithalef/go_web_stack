package config

import "errors"

type Config struct {
	MONGO_HOST     string `env:"MONGO_HOST"`
	MONGO_PORT     string `env:"MONGO_PORT"`
	MONGO_DATABASE string `env:"MONGO_DATABASE"`
	APP_PORT       string `env:"APP_PORT" envDefault:"8888"`
}

func (c Config) Validate() error {
	if len(c.MONGO_HOST) < 1 {
		return errors.New("environment variable MONGO_HOST is not set")
	}
	if len(c.MONGO_PORT) < 1 {
		return errors.New("environment variable MONGO_PORT is not set")
	}
	if len(c.MONGO_DATABASE) < 1 {
		return errors.New("environment variable MONGO_DATABASE is not set")
	}
	if len(c.APP_PORT) < 1 {
		return errors.New("environment variable APP_PORT is not set")
	}
	return nil
}
