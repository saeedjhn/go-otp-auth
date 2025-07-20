package models

import (
	"time"

	"github.com/saeedjhn/go-otp-auth/internal/types"
)

type User struct {
	ID        types.ID
	Name      string
	Mobile    string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
