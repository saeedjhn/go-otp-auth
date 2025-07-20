package claim

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-otp-auth/configs"
	"github.com/saeedjhn/go-otp-auth/internal/service/authentication"
)

func GetClaimsFromEchoContext(c echo.Context) *authentication.Claims {
	return c.Get(configs.AuthMiddlewareContextKey).(*authentication.Claims) //nolint:errcheck // nothing
}
