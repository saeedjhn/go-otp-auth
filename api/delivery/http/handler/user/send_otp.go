package user

import (
	"net/http"

	userdto "github.com/saeedjhn/go-otp-auth/internal/dto/user"

	"github.com/labstack/echo/v4"

	"github.com/saeedjhn/go-otp-auth/internal/models"

	"github.com/saeedjhn/go-otp-auth/pkg/bind"
	"github.com/saeedjhn/go-otp-auth/pkg/msg"

	"github.com/saeedjhn/go-otp-auth/pkg/httpstatus"
	"github.com/saeedjhn/go-otp-auth/pkg/richerror"
)

// SendOTP godoc.
// @Summary Send OTP to mobile
// @Description This endpoint sends a one-time password (OTP) to the provided mobile number.
// @Tags Users
// @Accept json
// @Produce json
// @Param Request body userdto.SendOTPRequest true "Mobile number to send OTP to"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request - Invalid input or binding error"
// @Failure 422 {object} models.ErrorResponse "Validation Failed"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /v1/users/send-otp [post]
func (h Handler) SendOTP(c echo.Context) error {
	req := userdto.SendOTPRequest{}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.NewErrorResponse(msg.ErrMsg400BadRequest, bind.CheckErrorFromBind(err).Error()).
				WithMeta(map[string]interface{}{"request": req}),
		)
	}

	resp, err := h.userSvc.SendOTP(c.Request().Context(), req)
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

	return c.JSON(http.StatusOK, models.NewSuccessResponse(msg.MsgOTPGenerated, resp))
}
