package main

import (
	"encoding/json"
	"fmt"
	"github.com/amithnair91/go_web_stack/load_test/app/config"
	"github.com/caarlos0/env"
	"strconv"
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


	metricsJson, _ := json.Marshal(metrics)
	printStat("Metrices", string(metricsJson))

	printStat("Throughput", strconv.FormatFloat(metrics.Throughput, 'E', -1, 64))

	printStats("Failures",metrics.Errors)
	printStat("Max-Latency", metrics.Latencies.Max.String())
	printStat("Min-Latency", metrics.Latencies.Min.String())
	printStat("Mean-Latency", metrics.Latencies.Mean.String())

	printStat("99th percentile: ", metrics.Latencies.P99.String())
}

func printStat(statName string, stat string) {
	fmt.Println(fmt.Sprintf("##############%s################", statName))
	fmt.Println(stat)
}

func printStats(statName string, stat []string) {
	fmt.Println(fmt.Sprintf("##############%s################", statName))
	fmt.Println(stat)
}

func itemCreationTarget(config config.Config) vegeta.Targeter {
	return vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    fmt.Sprintf("%s/%s", config.APP_BASE_URL, "item"),
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
