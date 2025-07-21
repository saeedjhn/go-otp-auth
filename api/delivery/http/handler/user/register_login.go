package user

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/saeedjhn/go-otp-auth/internal/models"

	userdto "github.com/saeedjhn/go-otp-auth/internal/dto/user"

	"github.com/saeedjhn/go-otp-auth/pkg/bind"
	"github.com/saeedjhn/go-otp-auth/pkg/httpstatus"
	"github.com/saeedjhn/go-otp-auth/pkg/msg"
	"github.com/saeedjhn/go-otp-auth/pkg/richerror"
)

// RegisterOrLogin godoc.
// @Summary Register or login a user
// @Description If the user does not exist, a new account will be created. If the user already exists,
// it will log them in.
// @Tags Auth
// @Accept json
// @Produce json
// @Param Request body userdto.RegisterOrLoginRequest true "User credentials for register or login"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request - Invalid input or binding error"
// @Failure 403 {object} models.ErrorResponse "Forbidden - Access denied"
// @Failure 422 {object} models.ErrorResponse "Validation Failed"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /v1/users/auth/register-or-login [post]
func (h Handler) RegisterOrLogin(c echo.Context) error {
	req := userdto.RegisterOrLoginRequest{}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.NewErrorResponse(msg.ErrMsg400BadRequest, bind.CheckErrorFromBind(err).Error()).
				WithMeta(map[string]interface{}{"request": req}),
		)
	}

	resp, err := h.userSvc.RegisterOrLogin(c.Request().Context(), req)
	if err != nil {
		richErr := richerror.Analysis(err)
		code := httpstatus.MapkindToHTTPStatusCode(richErr.Kind())

		if resp.FieldErrors != nil {
			return echo.NewHTTPError(
				code,
				models.NewErrorResponse(richErr.Message(), resp.FieldErrors).WithMeta(richErr.Meta()),
			)
		}

		return echo.NewHTTPError(
			code,
			models.NewErrorResponse(richErr.Message(), richErr.Error()).WithMeta(richErr.Meta()),
		)
	}

	return c.JSON(http.StatusOK, models.NewSuccessResponse(msg.MsgRegisterOrLoginSuccess, resp))
}
