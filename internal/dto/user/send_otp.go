package user

import "time"

type SendOTPRequest struct {
	Mobile string `json:"mobile" example:"09198829528"`
}
type SendOTPResponse struct {
	Mobile      string            `json:"mobile" example:"09198829528"`
	Code        string            `json:"code"`
	ExpireTime  time.Duration     `json:"expire_time"`
	FieldErrors map[string]string `json:"field_errors,omitempty" example:"{}"`
}
