package v1

import (
	"errors"

	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
	inventoryV1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
	"github.com/brianvoe/gofakeit/v7"
)

func (s *APISuite) TestGetPartSuccess() {
	var (
		uuid = "1"
		name = gofakeit.Name()

		modelPart = model.Part{
			Uuid: uuid,
			Name: name,
		}
		request = &inventoryV1.GetPartRequest{Uuid: uuid}
	)

	s.inventoryService.On("GetPart", s.ctx, uuid).Return(modelPart, nil)

	res, err := s.inventoryApi.GetPart(s.ctx, request)

	s.Require().NoError(err)
	s.Require().NotNil(res.Part)
	s.Require().Equal(res.Part.Uuid, uuid)
	s.Require().Equal(res.Part.Name, modelPart.Name)
}

func (s *APISuite) TestGetPartNotFound() {
	var (
		uuid = "1"
		name = gofakeit.Name()

		modelPart = model.Part{
			Uuid: uuid,
			Name: name,
		}
		request = &inventoryV1.GetPartRequest{Uuid: uuid}
	)

	s.inventoryService.On("GetPart", s.ctx, uuid).Return(modelPart, model.ErrPartNotFound)

	res, err := s.inventoryApi.GetPart(s.ctx, request)

	s.Require().Error(err)
	s.Require().Nil(res.Part)
	s.Require().True(errors.Is(err, model.ErrPartNotFound))
}

func (s *APISuite) TestGetPartError() {
	var (
		uuid = "1"
		name = gofakeit.Name()

		modelPart = model.Part{
			Uuid: uuid,
			Name: name,
		}
		request = &inventoryV1.GetPartRequest{Uuid: uuid}
	)

	s.inventoryService.On("GetPart", s.ctx, uuid).Return(modelPart, gofakeit.Error())

	res, err := s.inventoryApi.GetPart(s.ctx, request)

	s.Require().Error(err)
	s.Require().Nil(res.Part)
	s.Require().False(errors.Is(err, model.ErrPartNotFound))
}
