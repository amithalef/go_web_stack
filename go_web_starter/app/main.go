package main

import (
	"fmt"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/config"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/http/handlers"
	"net/http"

	"github.com/caarlos0/env"
)

func main() {
	config := config.Config{}
	env.Parse(&config)
	config.Validate()

	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/health-check", handlers.HealthCheckHandler)
	fmt.Println(fmt.Sprintf("Listening on port %s", config.APP_PORT))
	http.ListenAndServe(config.APP_PORT, nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
