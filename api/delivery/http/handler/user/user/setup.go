package user

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-otp-auth/internal/bootstrap"
	usermysql "github.com/saeedjhn/go-otp-auth/internal/repository/mysql/user"
	authservice "github.com/saeedjhn/go-otp-auth/internal/service/authentication"
	userservice "github.com/saeedjhn/go-otp-auth/internal/service/user"
	uservalidator "github.com/saeedjhn/go-otp-auth/internal/validator/user"
)

func Setup(app *bootstrap.Application, e *echo.Echo) {
	// Dependencies
	repo := usermysql.New(app.MySQL)

	vld := uservalidator.New(app.Config.Application.EntropyPassword)

	authSvc := authservice.New(app.Config.Auth)
	userSvc := userservice.New(app.Config, authSvc, vld, repo)

	handler := New(authSvc, userSvc)

	// Way 1
	handler.SetRoutes(e)
}
