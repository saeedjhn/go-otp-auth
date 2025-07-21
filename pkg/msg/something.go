package msg

const (
	ErrMsgSomethingWentWrong   = "something went wrong"
	ErrMsgCantPrepareStatement = "failed to prepare SQL statement"
	ErrMsgMobileIsNotValid     = "mobile is not valid"
	ErrMsgInvalidInput         = "invalid input"
	ErrMsgOTPCodeIsNotValid    = "code is not valid"

	ErrMsgDBRecordNotFound      = "record not found"
	ErrMsgDBCantScanQueryResult = "can't scan query result"

	MsgOTPGenerated           = "OTP generated successfully"
	MsgRegisterOrLoginSuccess = "Register or login successful"
)
