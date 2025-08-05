package main

import (
	"go-cqrs-api/api"
	"go-cqrs-api/config"
	"go-cqrs-api/infrastructure"
	"log"
	"net/http"
)

func main() {
	config.Load()

	infrastructure.InitDB()
	//infrastructure.InitRedis()
	//infrastructure.InitRabbitMQ()

	//go infrastructure.StartEventSubscribers()

	r := api.SetupRouter()
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
