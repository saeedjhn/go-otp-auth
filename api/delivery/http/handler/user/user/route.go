package user

import "github.com/labstack/echo/v4"

func (h Handler) SetRoutes(e *echo.Echo) {
	group := e.Group("/users")
	{
		// publicG := group.Group("")
		// {
		// publicG.POST("/refresh-token", h.RefreshToken)
		// }

		authG := group.Group("/auth")
		{
			authG.POST("/register", h.Register)
			authG.POST("/login", h.Login)
		}

		// protectedG := group.Group("")
		// protectedG.Use(mymiddleware.Authentication(h.authSvc))
		// {
		// 	protectedG.GET("/profile", h.Profile)
		// }
	}
}
