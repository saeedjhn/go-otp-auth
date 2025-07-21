package authentication

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/saeedjhn/go-otp-auth/internal/types"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID types.ID `json:"user_id"`
}

// func (c Claims) Valid() error {
//	return c.RegisteredClaims.Valid()
// }
