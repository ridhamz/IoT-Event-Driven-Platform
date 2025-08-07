package main

import (
	"go-cqrs-api/api"
	"go-cqrs-api/config"
	"go-cqrs-api/infrastructure"
	"go-cqrs-api/logger"
	"net/http"
)

func main() {
	config.Load()
	logger.Init()
	infrastructure.InitDB()
	infrastructure.InitSQS()
	//infrastructure.InitRabbitMQ()

	//go infrastructure.StartEventSubscribers()
	logger.Log.Info("App started")

	r := api.InitRouter()
	http.ListenAndServe(":8080", r)
}

// Protected routes
// r.Group(func(r chi.Router) {
// 	r.Use(middleware.AuthMiddleware)
// 	r.Get("/me", GetUserProfile)
// })
