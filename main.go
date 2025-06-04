package main

import (
	"github.com/cendolin/backend-auth-service/config"
	"github.com/cendolin/backend-auth-service/database"
	"github.com/cendolin/backend-auth-service/server"
)

func main() {
	config := config.NewConfig()

	db := database.NewDatabase(config)
	db.Migrate()

	server := server.NewServer(config, db)
	if err := server.Start(); err != nil {
		panic(err)
	}
}
