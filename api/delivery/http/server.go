package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/saeedjhn/go-otp-auth/api/delivery/http/handler"

	"github.com/labstack/echo/v4/middleware"
	mymiddleware "github.com/saeedjhn/go-otp-auth/api/delivery/http/middleware"
	"github.com/saeedjhn/go-otp-auth/internal/bootstrap"
)

type Server struct {
	App    *bootstrap.Application
	Router *echo.Echo
}

func New(
	app *bootstrap.Application,
) Server {
	return Server{
		App:    app,
		Router: echo.New(),
	}
}

func (s Server) Run() error {
	s.Router.Debug = s.App.Config.Application.Debug

	s.Router.Use(middleware.Recover())
	s.Router.Use(middleware.RequestID())
	s.Router.Use(mymiddleware.Timeout(s.App.Config.HTTPServer.Timeout))

	handler.Setup(s.App, s.Router)

	address := fmt.Sprintf(":%s", s.App.Config.HTTPServer.Port)

	if err := s.Router.Start(address); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
