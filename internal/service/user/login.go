package user

import (
	"context"

	"github.com/saeedjhn/go-otp-auth/internal/models"

	"github.com/saeedjhn/go-otp-auth/internal/dto/user"
	"github.com/saeedjhn/go-otp-auth/pkg/msg"
	"github.com/saeedjhn/go-otp-auth/pkg/richerror"
)

func (i Service) Login(ctx context.Context, req user.LoginRequest) (user.LoginResponse, error) {
	if fieldsErrs, err := i.vld.ValidateLoginRequest(req); err != nil {
		return user.LoginResponse{FieldErrors: fieldsErrs}, err
	}

	u, err := i.repository.GetByMobile(ctx, req.Mobile)
	if err != nil {
		return user.LoginResponse{}, err
	}

	err = CompareHash(u.Password, req.Password)
	if err != nil {
		return user.LoginResponse{}, richerror.New(_opUserServiceLogin).WithErr(err).
			WithMessage(errMsgIncorrectPassword).
			WithKind(richerror.KindStatusBadRequest)
	}

	authenticable := models.Authenticate{ID: u.ID}

	accessToken, err := i.authSvc.CreateAccessToken(authenticable)
	if err != nil {
		return user.LoginResponse{}, richerror.New(_opUserServiceLogin).WithErr(err).
			WithMessage(msg.ErrorMsg500InternalServerError).
			WithKind(richerror.KindStatusInternalServerError)
	}

	refreshToken, err := i.authSvc.CreateRefreshToken(authenticable)
	if err != nil {
		return user.LoginResponse{}, richerror.New(_opUserServiceLogin).WithErr(err).
			WithMessage(msg.ErrorMsg500InternalServerError).
			WithKind(richerror.KindStatusInternalServerError)
	}

	return user.LoginResponse{
		UserInfo: user.Info{
			ID:        u.ID,
			Name:      u.Name,
			Mobile:    u.Mobile,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		},
		Tokens: user.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}
