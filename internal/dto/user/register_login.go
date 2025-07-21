package user

type RegisterOrLoginRequest struct {
	Mobile string `json:"mobile" example:"09198829528"`
	Code   string `json:"code" example:"12345"`
}

type RegisterOrLoginResponse struct {
	UserInfo    Info              `json:"user"`
	Tokens      Tokens            `json:"tokens"`
	FieldErrors map[string]string `json:"field_errors,omitempty"`
}
