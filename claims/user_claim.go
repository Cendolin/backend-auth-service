package claims

import "github.com/golang-jwt/jwt/v5"

type UserClaim struct {
	jwt.RegisteredClaims

	Username string `json:"username"`
	Id       string `json:"user_id"`
	Country  string `json:"country"`
	Email    string `json:"email"`
}
