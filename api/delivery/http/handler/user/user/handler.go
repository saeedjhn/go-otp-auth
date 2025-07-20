package user

import (
	authservice "github.com/saeedjhn/go-otp-auth/internal/service/authentication"
	userservice "github.com/saeedjhn/go-otp-auth/internal/service/user"
)

type Handler struct {
	authSvc authservice.Service
	userSvc userservice.Service
}

func New(authSvc authservice.Service, userSvc userservice.Service) Handler {
	return Handler{authSvc: authSvc, userSvc: userSvc}
}
