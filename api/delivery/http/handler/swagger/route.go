package swagger

import (
	"github.com/labstack/echo/v4"
	// Import Swagger docs for side effects only (required by swagger).
	_ "github.com/saeedjhn/go-otp-auth/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	group := e.Group("/v1/swagger")
	{
		group.GET("/*", echoSwagger.WrapHandler)
	}
}
