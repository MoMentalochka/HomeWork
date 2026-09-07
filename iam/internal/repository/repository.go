package repository

import (
	"context"

	"github.com/MoMentalochka/HomeWork/iam/internal/model"
	iamV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/iam/v1"
)

type IamUserRepository interface {
	Get(ctx context.Context, userUuid string) (*iamV1.GetUserResponse, error)
	Create(ctx context.Context, userDTO *iamV1.RegisterRequest) (*iamV1.RegisterResponse, error)
}

type IamSessionRepository interface {
	Get(ctx context.Context, sessionUuid string) (*model.Session, error)
	Create(ctx context.Context, sessionDTO *model.User) (*iamV1.LoginResponse, error)
}
