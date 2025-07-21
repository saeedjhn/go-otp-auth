package user

import (
	"errors"
	"regexp"

	"github.com/saeedjhn/go-otp-auth/configs"

	"github.com/saeedjhn/go-otp-auth/pkg/msg"

	userdto "github.com/saeedjhn/go-otp-auth/internal/dto/user"

	"github.com/saeedjhn/go-otp-auth/pkg/richerror"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateRegisterOrLoginRequest(req userdto.RegisterOrLoginRequest) (map[string]string, error) {
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Mobile,
			validation.Required,
			validation.Length(_mobileMinLen, _mobileMaxLen),
			validation.Match(regexp.MustCompile(configs.MobileRegex)).Error(msg.ErrMsgMobileIsNotValid),
		),

		validation.Field(&req.Code,
			validation.Required,
			validation.Length(configs.OTPLen, configs.OTPLen)),
	); err != nil {
		var fieldErrors = make(map[string]string)

		var errV validation.Errors
		ok := errors.As(err, &errV)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}

		return fieldErrors, richerror.New(_opUserValidatorValidateLoginRequest).
			WithErr(err).
			WithMessage(msg.ErrMsgInvalidInput).
			WithKind(richerror.KindStatusUnprocessableEntity).
			WithMeta(map[string]interface{}{"req": req})
	}

	return nil, nil //nolint:nilnil // return both the `nil` error and invalid value: use a sentinel error instead (nilnil)
}
