package controllers

import (
	"github.com/cendolin/backend-auth-service/dtos"
	"github.com/cendolin/backend-auth-service/models"
	checker "github.com/cinar/checker/v2"
	"github.com/gofiber/fiber/v3"
	"github.com/matthewhartstonge/argon2"
)

func (t *Controllers) RegisterController(ctx fiber.Ctx) error {
	payload := &dtos.RegisterDto{}
	if err := ctx.Bind().Body(payload); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors, valid := checker.CheckStruct(payload)
	if !valid {
		return ctx.Status(400).JSON(fiber.Map{
			"errors":  errors,
			"message": "Validation Error",
		})
	}

	count := int64(0)
	if err := t.DB.DB.Model(&models.User{}).Where("username = ?", payload.Username).Or("email = ?", payload.Email).Where("suspended_reason IS NOT NULL").Count(&count).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if count > 0 {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "username or email is already taken",
		})
	}

	argon := argon2.DefaultConfig()
	pw, err := argon.HashEncoded([]byte(payload.Password))
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user_payload := &models.User{
		Email:        payload.Email,
		Username:     payload.Username,
		PasswordHash: string(pw),
		Biography:    payload.Biography,
		Country:      payload.Country,
	}

	if err := t.DB.DB.Model(&models.User{}).Create(user_payload).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"user": user_payload,
	})
}
