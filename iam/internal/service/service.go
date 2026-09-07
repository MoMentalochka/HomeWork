package service

import (
	"context"

	iamV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/iam/v1"
)

type IamAuthService interface {
	Login(ctx context.Context, request *iamV1.LoginRequest) (*iamV1.LoginResponse, error)
	Whoami(ctx context.Context, request *iamV1.WhoamiRequest) (*iamV1.WhoamiResponse, error)
}

type IamUserService interface {
	Register(ctx context.Context, request *iamV1.RegisterRequest) (*iamV1.RegisterResponse, error)
	GetUser(ctx context.Context, request *iamV1.GetUserRequest) (*iamV1.GetUserResponse, error)
}
