package main

import (
	"github.com/cendolin/backend-auth-service/config"
	"github.com/cendolin/backend-auth-service/database"
	"github.com/cendolin/backend-auth-service/rabbit"
	"github.com/cendolin/backend-auth-service/server"
)

func main() {
	config := config.NewConfig()

	// Database
	db := database.NewDatabase(config)
	db.Migrate()

	// RabbitMQ
	rab := rabbit.NewRabbit(config.RabbitMQUrl)
	rab.Init()

	defer rab.Close()

	server := server.NewServer(config, db, rab)
	if err := server.Start(); err != nil {
		panic(err)
	}
}
