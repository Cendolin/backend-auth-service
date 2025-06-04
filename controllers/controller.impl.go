package controllers

import (
	"github.com/cendolin/backend-auth-service/config"
	"github.com/cendolin/backend-auth-service/database"
)

type Controllers struct {
	DB     *database.Database
	Config *config.Config
}

func NewControllers(db *database.Database, config *config.Config) *Controllers {
	return &Controllers{
		DB:     db,
		Config: config,
	}
}
