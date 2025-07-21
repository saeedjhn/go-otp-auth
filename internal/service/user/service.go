package user

import (
	"context"
	"time"

	"github.com/saeedjhn/go-otp-auth/internal/contract"

	"github.com/saeedjhn/go-otp-auth/internal/service/authentication"

	userdto "github.com/saeedjhn/go-otp-auth/internal/dto/user"

	"github.com/saeedjhn/go-otp-auth/internal/models"

	"github.com/saeedjhn/go-otp-auth/configs"
)

//go:generate mockery --name AuthInteractor
type AuthService interface {
	CreateAccessToken(req models.Authenticate) (string, error)
	CreateRefreshToken(req models.Authenticate) (string, error)
	ParseToken(secret, requestToken string) (*authentication.Claims, error)
}

//go:generate mockery --name Validator
type Validator interface {
	ValidateSendOTPRequest(req userdto.SendOTPRequest) (map[string]string, error)
	ValidateRegisterOrLoginRequest(req userdto.RegisterOrLoginRequest) (map[string]string, error)
}

//go:generate mockery --name Cache
type Cache interface {
	SetByMobile(ctx context.Context, mobile string, code string, expireTime time.Duration) error
	GetByMobile(ctx context.Context, mobile string) (string, error)
	DelByMobile(ctx context.Context, mobile string) (bool, error)
}

//go:generate mockery --name Repository
type Repository interface {
	Create(ctx context.Context, u models.User) (models.User, error)
	ExistsByMobile(ctx context.Context, mobile string) (bool, error)
	GetByMobile(ctx context.Context, mobile string) (models.User, error)
}

type Service struct {
	cfg        *configs.Config
	logger     contract.Logger
	authSvc    AuthService
	vld        Validator
	cache      Cache
	repository Repository
}

func New(
	cfg *configs.Config,
	logger contract.Logger,
	authSvc AuthService,
	vld Validator,
	cache Cache,
	repository Repository,
) Service {
	return Service{
		cfg:        cfg,
		logger:     logger,
		vld:        vld,
		authSvc:    authSvc,
		cache:      cache,
		repository: repository,
	}
}
