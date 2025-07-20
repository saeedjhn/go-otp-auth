package user

import (
	"context"

	"github.com/saeedjhn/go-otp-auth/internal/service/authentication"

	"github.com/saeedjhn/go-otp-auth/pkg/security/bcrypt"

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
	ValidateRegisterRequest(req userdto.RegisterRequest) (map[string]string, error)
	ValidateLoginRequest(req userdto.LoginRequest) (map[string]string, error)
}

//go:generate mockery --name Repository
type Repository interface {
	Create(ctx context.Context, u models.User) (models.User, error)
	IsExistsByMobile(ctx context.Context, mobile string) (bool, error)
	IsExistsByEmail(ctx context.Context, email string) (bool, error)
	GetByMobile(ctx context.Context, mobile string) (models.User, error)
	GetByID(ctx context.Context, id uint64) (models.User, error)
}

type Service struct {
	cfg        *configs.Config
	authSvc    AuthService
	vld        Validator
	repository Repository
}

// var _ userhandler.Service = (Service)(nil) // Commented, because it happens import cycle.

func New(
	cfg *configs.Config,
	authSvc AuthService,
	vld Validator,
	repository Repository,
) Service {
	return Service{
		cfg:        cfg,
		vld:        vld,
		authSvc:    authSvc,
		repository: repository,
	}
}

func GenerateHash(password string) (string, error) {
	return bcrypt.Generate(password, bcrypt.Cost(configs.BcryptCost))
}

func CompareHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndSTR(hashedPassword, password)
}
