package v1

import (
	"github.com/MoMentalochka/HomeWork/iam/internal/service"
	iamV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/iam/v1"
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
