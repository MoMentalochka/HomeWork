package v1

import (
	"context"

	"github.com/MoMentalochka/HomeWork/iam/internal/service"
	"github.com/MoMentalochka/HomeWork/platform/pkg/logger"
	iamV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/iam/v1"
	"go.uber.org/zap"
)

type api struct {
	iamV1.UnimplementedIamServiceServer
	iamService service.IamService
}

func NewApi(iamService service.IamService) *api {
	return &api{
		iamService: iamService,
	}
}

func (a *api) Login(ctx context.Context, request *iamV1.LoginRequest) (*iamV1.LoginResponse, error) {
	logger.Info(ctx, "Login", zap.String("login", request.Login), zap.String("pass", request.Password))
	return &iamV1.LoginResponse{SessionUuid: "1"}, nil
}

func (a *api) Whoami(ctx context.Context, request *iamV1.WhoamiRequest) (*iamV1.WhoamiResponse, error) {
	logger.Info(ctx, "Whoami", zap.String("SessionUuid", request.SessionUuid))
	return &iamV1.WhoamiResponse{UserUuid: "1"}, nil
}

func (a *api) Register(ctx context.Context, request *iamV1.RegisterRequest) (*iamV1.RegisterResponse, error) {
	logger.Info(ctx, "Register", zap.String("login", request.Login), zap.String("pass", request.Password), zap.String("email", request.Email))
	return &iamV1.RegisterResponse{UserUuid: "1"}, nil
}

func (a *api) GetUser(ctx context.Context, request *iamV1.GetUserRequest) (*iamV1.GetUserResponse, error) {
	logger.Info(ctx, "GetUser", zap.String("UserUuid", request.UserUuid))
	return &iamV1.GetUserResponse{UserUuid: "1"}, nil
}
