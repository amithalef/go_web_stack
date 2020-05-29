package config

import "errors"

type Config struct {
	APP_BASE_URL               string `env:"APP_BASE_URL"`
	ATTACK_RATE_PER_SECOND     int    `env:"ATTACK_RATE_PER_SECOND" envDefault:"100"`
	ATTACK_DURATION_IN_SECONDS int    `env:"ATTACK_DURATION_IN_SECONDS" envDefault:"10"`
}

func (c Config) Validate() error {
	if len(c.APP_BASE_URL) < 1 {
		return errors.New("environment variable APP_BASE_URL is not set")
	}
	return nil
}
