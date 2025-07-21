package user

import (
	"errors"
	"regexp"

	"github.com/saeedjhn/go-otp-auth/configs"
	userdto "github.com/saeedjhn/go-otp-auth/internal/dto/user"
	"github.com/saeedjhn/go-otp-auth/pkg/msg"
	"github.com/saeedjhn/go-otp-auth/pkg/richerror"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateSendOTPRequest(req userdto.SendOTPRequest) (map[string]string, error) {
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Mobile,
			validation.Required,
			validation.Length(_mobileMinLen, _mobileMaxLen),
			validation.Match(regexp.MustCompile(configs.MobileRegex)).Error(msg.ErrMsgMobileIsNotValid),
		)); err != nil {
		fieldErrors := make(map[string]string)

		vErr := validation.Errors{}
		if errors.As(err, &vErr) {
			for key, value := range vErr {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}

		return fieldErrors, richerror.New(_opUserValidatorValidateSendOTPRequest).
			WithErr(err).
			WithMessage(msg.ErrMsgInvalidInput).
			WithKind(richerror.KindStatusUnprocessableEntity).
			WithMeta(map[string]interface{}{"req": req})
	}

	return nil, nil //nolint:nilnil // return both the `nil` error and invalid value: use a sentinel error instead (nilnil)
}
