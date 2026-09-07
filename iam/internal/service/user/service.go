package user

import (
	"context"

	"github.com/MoMentalochka/HomeWork/iam/internal/repository"
	"github.com/MoMentalochka/HomeWork/platform/pkg/logger"
	iamV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/iam/v1"
	"go.uber.org/zap"
)

type iamUserService struct {
	userRepository repository.IamUserRepository
}

func NewIamUserService(userRepository repository.IamUserRepository) *iamUserService {
	return &iamUserService{
		userRepository,
	}
}

func (a *iamUserService) Register(ctx context.Context, request *iamV1.RegisterRequest) (*iamV1.RegisterResponse, error) {
	logger.Info(ctx, "Register", zap.String("login", request.Login), zap.String("pass", request.Password), zap.String("email", request.Email))
	return &iamV1.RegisterResponse{UserUuid: "1"}, nil
}

func (a *iamUserService) GetUser(ctx context.Context, request *iamV1.GetUserRequest) (*iamV1.GetUserResponse, error) {
	logger.Info(ctx, "GetUser", zap.String("UserUuid", request.UserUuid))
	return &iamV1.GetUserResponse{UserUuid: "1"}, nil
}
