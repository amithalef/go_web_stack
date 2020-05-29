package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"

	"github.com/caarlos0/env"

	"github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/config"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/http/handlers"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/mongo_storage"
)

func main() {
	config := initializeConfig()
	database := connect(config)
	itemHandler := itemHandler(database)

	http.HandleFunc("/health-check", handlers.HealthCheckHandler)
	http.HandleFunc("/item", itemHandler.CreateItemHandler)

	fmt.Println(fmt.Sprintf("Listening on port %s", config.APP_PORT))
	http.ListenAndServe(fmt.Sprintf(":%s",config.APP_PORT), nil)
}

func itemHandler(database *mongo.Database) handlers.ItemHttpHandler {
	itemStorage := mongo_storage.NewMongoItemStorage(database)
	itemHandler := handlers.ItemHttpHandler{ItemStorage: itemStorage}
	return itemHandler
}

func connect(config config.Config) *mongo.Database {
	database, err := mongo_storage.Connect(config.MONGO_HOST, config.MONGO_PORT, config.MONGO_DATABASE)
	if err != nil {
		panic(fmt.Sprintf("application cannot connect to mongodb with config %#v", err))
		os.Exit(1)
	}
	return database
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
