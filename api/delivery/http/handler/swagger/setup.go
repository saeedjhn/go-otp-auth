package swagger

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-otp-auth/internal/bootstrap"
)

func Setup(_ *bootstrap.Application, e *echo.Echo) {
	handler := New()
	handler.SetRoutes(e)
}
