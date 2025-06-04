package claims

import "github.com/golang-jwt/jwt/v5"

type UserClaim struct {
	jwt.RegisteredClaims

	Username string `json:"username"`
	Id       string `json:"user_id"`
}
