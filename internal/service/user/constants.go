package user

const (
	_opUserServiceSendOTP         = "userservice_SendOTP"
	_opUserServiceRegisterOrLogin = "userservice_RegisterOrLogin"

	errMsgMobileIsNotUnique            = "mobile is not unique"
	errMsgEmailIsNotUnique             = "email address is not unique"
	errMsgFailedToGeneratePasswordHash = "failed to generate password hash "
	errMsgIncorrectPassword            = "the password is incorrect"
)
