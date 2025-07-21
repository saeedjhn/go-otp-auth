package handler

import (
	"github.com/labstack/echo/v4"
	swaggerhandler "github.com/saeedjhn/go-otp-auth/api/delivery/http/handler/swagger"
	userhandler "github.com/saeedjhn/go-otp-auth/api/delivery/http/handler/user"
	"github.com/saeedjhn/go-otp-auth/internal/bootstrap"
)

func Setup(app *bootstrap.Application, e *echo.Echo) {
	userhandler.Setup(app, e)
	swaggerhandler.Setup(app, e)
}
