package controllers

import (
	"github.com/cendolin/backend-auth-service/config"
	"github.com/cendolin/backend-auth-service/database"
	"github.com/cendolin/backend-auth-service/rabbit"
)

type Controllers struct {
	DB     *database.Database
	Config *config.Config
	Rabbit *rabbit.Rabbit
}

func NewControllers(db *database.Database, config *config.Config, rabbit *rabbit.Rabbit) *Controllers {
	return &Controllers{
		DB:     db,
		Config: config,
		Rabbit: rabbit,
	}
}
