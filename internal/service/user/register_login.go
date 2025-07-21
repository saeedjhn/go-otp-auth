package user

import (
	"context"

	"github.com/saeedjhn/go-otp-auth/internal/models"

	userdto "github.com/saeedjhn/go-otp-auth/internal/dto/user"
	"github.com/saeedjhn/go-otp-auth/pkg/msg"
	"github.com/saeedjhn/go-otp-auth/pkg/richerror"
)

func (s Service) RegisterOrLogin(
	ctx context.Context,
	req userdto.RegisterOrLoginRequest,
) (userdto.RegisterOrLoginResponse, error) {
	if fieldsErrs, vErr := s.vld.ValidateRegisterOrLoginRequest(req); vErr != nil {
		return userdto.RegisterOrLoginResponse{FieldErrors: fieldsErrs},
			richerror.New(_opUserServiceRegisterOrLogin).WithErr(vErr)
	}

	code, gErr := s.cache.GetByMobile(ctx, req.Mobile)
	if gErr != nil {
		return userdto.RegisterOrLoginResponse{}, richerror.New(_opUserServiceRegisterOrLogin).
			WithErr(gErr).
			WithKind(richerror.KindStatusInternalServerError)
	}
	if code == "" || code != req.Code {
		return userdto.RegisterOrLoginResponse{}, richerror.New(_opUserServiceRegisterOrLogin).
			WithMessage(msg.ErrMsgOTPCodeIsNotValid).
			WithKind(richerror.KindStatusForbidden)
	}

	_, dErr := s.cache.DelByMobile(ctx, req.Mobile)
	if dErr != nil {
		return userdto.RegisterOrLoginResponse{}, richerror.New(_opUserServiceRegisterOrLogin).
			WithErr(dErr).
			WithKind(richerror.KindStatusInternalServerError)
	}

	var user models.User
	isExist, rErr := s.repository.ExistsByMobile(ctx, req.Mobile)
	if rErr != nil {
		return userdto.RegisterOrLoginResponse{}, richerror.New(_opUserServiceRegisterOrLogin).
			WithErr(rErr).
			WithKind(richerror.KindStatusInternalServerError)
	}

	if !isExist {
		userCreated, err := s.repository.Create(ctx, models.User{Mobile: req.Mobile})
		if err != nil {
			return userdto.RegisterOrLoginResponse{}, richerror.New(_opUserServiceRegisterOrLogin).
				WithErr(err).
				WithKind(richerror.KindStatusInternalServerError)
		}

		user = userCreated
	} else {
		getUser, err := s.repository.GetByMobile(ctx, req.Mobile)
		if err != nil {
			return userdto.RegisterOrLoginResponse{}, err
		}

		user = getUser
	}

	authenticate := models.Authenticate{ID: user.ID}

	accessToken, aErr := s.authSvc.CreateAccessToken(authenticate)
	if aErr != nil {
		return userdto.RegisterOrLoginResponse{}, richerror.New(_opUserServiceRegisterOrLogin).
			WithErr(aErr).
			WithMessage(msg.ErrorMsg500InternalServerError).
			WithKind(richerror.KindStatusInternalServerError)
	}

	refreshToken, rErr := s.authSvc.CreateRefreshToken(authenticate)
	if rErr != nil {
		return userdto.RegisterOrLoginResponse{}, richerror.New(_opUserServiceRegisterOrLogin).
			WithErr(rErr).
			WithMessage(msg.ErrorMsg500InternalServerError).
			WithKind(richerror.KindStatusInternalServerError)
	}

	return userdto.RegisterOrLoginResponse{
		UserInfo: userdto.Info{
			ID:        user.ID,
			Mobile:    user.Mobile,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Tokens: userdto.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}
