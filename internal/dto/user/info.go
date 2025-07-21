package user

import (
	"time"

	"github.com/saeedjhn/go-otp-auth/internal/types"
)

type Info struct {
	ID        types.ID  `json:"id"`
	Mobile    string    `json:"mobile"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
