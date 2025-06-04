package controllers

import (
	"time"

	"github.com/cendolin/backend-auth-service/claims"
	"github.com/cendolin/backend-auth-service/dtos"
	"github.com/cendolin/backend-auth-service/models"
	checker "github.com/cinar/checker/v2"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/matthewhartstonge/argon2"
)

func (t *Controllers) LoginController(ctx fiber.Ctx) error {
	payload := &dtos.LoginDto{}
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

	user := &models.User{}
	if err := t.DB.DB.Model(&models.User{}).Select("password_hash", "username", "email", "id", "verified").Where("username = ?", payload.Username).Or("email = ?", payload.Email).Where("suspended_reason IS NULL").Where("verified = true").First(user).Error; err != nil {
		return ctx.Status(401).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if ok, err := argon2.VerifyEncoded([]byte(payload.Password), []byte(user.PasswordHash)); err != nil {
		return ctx.Status(401).JSON(fiber.Map{
			"message": err.Error(),
		})
	} else if !ok {
		return ctx.Status(401).JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	claims := &claims.UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Cendolin",
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(2 * time.Hour)},
		},
		Username: user.Username,
		Id:       user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedToken, err := token.SignedString([]byte(t.Config.Jwt.Key))
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"data": map[string]string{
			"token":    signedToken,
			"user_id":  user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}
