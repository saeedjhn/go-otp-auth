package user

import (
	"context"
	"log"

	"github.com/saeedjhn/go-otp-auth/configs"
	userdto "github.com/saeedjhn/go-otp-auth/internal/dto/user"
	"github.com/saeedjhn/go-otp-auth/pkg/generator"
	"github.com/saeedjhn/go-otp-auth/pkg/msg"
	"github.com/saeedjhn/go-otp-auth/pkg/richerror"
)

func (s Service) SendOTP(ctx context.Context, req userdto.SendOTPRequest) (userdto.SendOTPResponse, error) {
	if fieldErrors, vErr := s.vld.ValidateSendOTPRequest(req); vErr != nil {
		return userdto.SendOTPResponse{FieldErrors: fieldErrors}, richerror.New(_opUserServiceSendOTP).WithErr(vErr)
	}

	// TODO: Optimization - Check cache first:
	// If an OTP code already exists for the given mobile number, return the existing code
	// to prevent generating a new one within the expiration time.

	newCode, gErr := generator.GenCode(configs.OTPLen, configs.OTPChars)
	if gErr != nil {
		return userdto.SendOTPResponse{}, richerror.New(_opUserServiceSendOTP).
			WithErr(gErr).
			WithMessage(msg.ErrorMsg500InternalServerError).
			WithKind(richerror.KindStatusInternalServerError)
	}

	cErr := s.cache.SetByMobile(ctx, req.Mobile, newCode, configs.OTPExpireTime)
	if cErr != nil {
		return userdto.SendOTPResponse{}, richerror.New(_opUserServiceSendOTP).
			WithErr(cErr).
			WithKind(richerror.KindStatusInternalServerError)
	}

	log.Printf("Generated OTP code: %s for mobile: %s\n", newCode, req.Mobile)

	return userdto.SendOTPResponse{
		Mobile:     req.Mobile,
		Code:       newCode,               // TODO - have to remove it in production
		ExpireTime: configs.OTPExpireTime, // TODO - have to remove it in production
	}, nil
}
