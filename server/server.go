package server

import (
	"fmt"

	"github.com/bytedance/sonic"
	"github.com/cendolin/backend-auth-service/config"
	"github.com/cendolin/backend-auth-service/controllers"
	"github.com/cendolin/backend-auth-service/database"
	"github.com/cendolin/backend-auth-service/rabbit"
	"github.com/gofiber/fiber/v3"

	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/idempotency"
)

func NewServer(config *config.Config, db *database.Database, rabbit *rabbit.Rabbit) *Server {
	app := fiber.New(fiber.Config{
		ServerHeader: "Cendolin",
		AppName:      "Cendolin Auth Service",
		TrustProxy:   true,
		JSONEncoder:  sonic.Marshal,
		JSONDecoder:  sonic.Unmarshal,
	})

	return &Server{
		App:    app,
		Config: config,
		DB:     db,
		Rabbit: rabbit,
	}
}

// Implementations

func (s *Server) init() {
	s.App.Use(cors.New())
	s.App.Use(helmet.New())
	s.App.Use(idempotency.New())

	controllers := controllers.NewControllers(s.DB, s.Config, s.Rabbit)

	s.App.Get("/", func(c fiber.Ctx) error {
		return c.SendStatus(200)
	})

	s.App.Post("/login", controllers.LoginController)
	s.App.Post("/register", controllers.RegisterController)
}

func (s *Server) Start() error {
	s.init() // Initializing routes

	formatted_addr := fmt.Sprintf("%s:%d", s.Config.Api.ListenHost, s.Config.Api.ListenPort)
	return s.App.Listen(formatted_addr)
}
