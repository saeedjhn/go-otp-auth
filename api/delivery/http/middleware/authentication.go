package middleware

import (
	mw "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-otp-auth/configs"
	"github.com/saeedjhn/go-otp-auth/internal/service/authentication"
)

func Authentication(
	authSvc authentication.Service,
) echo.MiddlewareFunc {
	return mw.WithConfig(mw.Config{
		ContextKey: configs.AuthMiddlewareContextKey,
		SigningKey: []byte(authSvc.Config.AccessTokenSecret),
		// TODO  - as sign method string to config
		SigningMethod: "HS256",
		ParseTokenFunc: func(_ echo.Context, auth string) (interface{}, error) {
			claims, err := authSvc.ParseToken(authSvc.Config.AccessTokenSecret, auth)
			if err != nil {
				return nil, err
			}

			return claims, nil
		},
	})
}

// const _lenValidAuthorizationKeyFromHeader = 2
//
// func Authentication(
//	authService *authservice.Service,
// ) echo.MiddlewareFunc {
//	return func(next echo.HandlerFunc) echo.HandlerFunc {
//		return func(c echo.Context) error {
//			authHeader := c.Request().Header.Get("Authorization")
//			t := strings.Split(authHeader, " ")
//
//			if len(t) == _lenValidAuthorizationKeyFromHeader {
//				authToken := t[1]
//				authorized, err := authService.IsAuthorized(
//					authToken,
//					authService.Config.AccessTokenSecret,
//				)
//				if authorized {
//					parseTokenResp, errParse := authService.ParseToken(userauthservicedto.ParseTokenRequest{
//						Secret: authService.Config.AccessTokenSecret,
//						Token:  authToken,
//					})
//					if errParse != nil {
//						return c.JSON(http.StatusUnauthorized, echo.Map{
//							"status":  false,
//							"message": message.ErrorMsg401UnAuthorized,
//							"errors":  nil,
//						})
//					}
//
//					claim.SetClaimsFromEchoContext(c, configs.AuthMiddlewareContextKey, parseTokenResp.Claims)
//
//					return next(c)
//				}
//
//				return c.JSON(http.StatusUnauthorized, echo.Map{
//					"status":  false,
//					"message": message.ErrorMsg401UnAuthorized,
//					"errors":  err.Error(),
//				})
//			}
//
//			return c.JSON(http.StatusUnauthorized, echo.Map{
//				"status":  false,
//				"message": message.ErrorMsg401UnAuthorized,
//				"errors":  nil,
//			})
//		}
//	}
// }
