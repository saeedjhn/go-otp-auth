package user

import "github.com/labstack/echo/v4"

func (h Handler) SetRoutes(e *echo.Echo) {
	group := e.Group("/v1/users")
	{
		pubGroup := group.Group("")
		{
			pubGroup.POST("/send-otp", h.SendOTP)
		}

		authGroup := group.Group("/auth")
		{
			authGroup.POST("/register-or-login", h.RegisterOrLogin)
		}
	}
}
