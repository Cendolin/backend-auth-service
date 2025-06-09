package server

import (
	"github.com/cendolin/backend-auth-service/config"
	"github.com/cendolin/backend-auth-service/database"
	"github.com/cendolin/backend-auth-service/rabbit"
	"github.com/gofiber/fiber/v3"
)

type Server struct {
	App    *fiber.App
	Config *config.Config
	DB     *database.Database
	Rabbit *rabbit.Rabbit
}
