package main

import (
	"fmt"
	"github.com/amithnair91/go_web_stack/load_test/app/config"
	"github.com/caarlos0/env"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	config := initializeConfig()
	rate := attackRatePerSecond(config)
	duration := attachDurationInSeconds(config)
	itemCreationTarget := itemCreationTarget(config)
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(itemCreationTarget, rate, duration, "Create Lots of Items!") {
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}

func itemCreationTarget(config config.Config) vegeta.Targeter {
	return vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    fmt.Sprintf("%s/%s", config.APP_BASE_URL, "/item"),
		Body: []byte(`{"Name":"bag"}`),
	})
}

func attackRatePerSecond(config config.Config) vegeta.Rate {
	return vegeta.Rate{Freq: config.ATTACK_RATE_PER_SECOND, Per: time.Second}
}

func attachDurationInSeconds(config config.Config) time.Duration {
	return time.Duration(config.ATTACK_DURATION_IN_SECONDS) * time.Second
}

func initializeConfig() config.Config {
	config := config.Config{}
	env.Parse(&config)
	err := config.Validate()
	if err != nil {
		panic(fmt.Sprintf("the supplied configuration is not valid %#v : %s", config, err.Error()))
	}
	fmt.Println(fmt.Sprintf("initializing application with config %#v", config))
	return config
}
