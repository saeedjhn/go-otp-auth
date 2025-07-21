package models

import (
	"time"

	"github.com/saeedjhn/go-otp-auth/internal/types"
)

type User struct {
	ID        types.ID
	Mobile    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
