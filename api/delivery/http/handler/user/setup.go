package user

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-otp-auth/internal/bootstrap"
	usermysql "github.com/saeedjhn/go-otp-auth/internal/repository/mysql/user"
	otpredis "github.com/saeedjhn/go-otp-auth/internal/repository/redis/otp"
	authservice "github.com/saeedjhn/go-otp-auth/internal/service/authentication"
	userservice "github.com/saeedjhn/go-otp-auth/internal/service/user"
	uservalidator "github.com/saeedjhn/go-otp-auth/internal/validator/user"
)

func Setup(app *bootstrap.Application, e *echo.Echo) {
	repo := usermysql.New(app.MySQL)
	redisRepo := otpredis.New(app.Redis)

	vld := uservalidator.New()

	authSvc := authservice.New(app.Config.Auth)
	userSvc := userservice.New(
		app.Config,
		app.Logger,
		authSvc,
		vld,
		redisRepo,
		repo,
	)

	handler := New(authSvc, userSvc)

	handler.SetRoutes(e)
}
