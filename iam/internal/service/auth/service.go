package auth

import (
	"context"
	"time"

	"github.com/MoMentalochka/HomeWork/iam/internal/repository"
	"github.com/MoMentalochka/HomeWork/platform/pkg/logger"
	iamV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/iam/v1"
	"go.uber.org/zap"
)

type iamAuthService struct {
	sessionRepository repository.IamSessionRepository
	sessionTTL        time.Duration
}

func NewIamAuthService(session repository.IamSessionRepository, sessionTTL time.Duration) *iamAuthService {
	return &iamAuthService{
		session,
		sessionTTL,
	}
}

func (a *iamAuthService) Login(ctx context.Context, request *iamV1.LoginRequest) (*iamV1.LoginResponse, error) {
	logger.Info(ctx, "Login", zap.String("login", request.Login), zap.String("pass", request.Password))
	return &iamV1.LoginResponse{SessionUuid: "1"}, nil
}

func (a *iamAuthService) Whoami(ctx context.Context, request *iamV1.WhoamiRequest) (*iamV1.WhoamiResponse, error) {
	logger.Info(ctx, "Whoami", zap.String("SessionUuid", request.SessionUuid))
	return &iamV1.WhoamiResponse{UserUuid: "1"}, nil
}
