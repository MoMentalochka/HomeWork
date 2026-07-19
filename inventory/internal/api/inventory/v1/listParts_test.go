package v1

import (
	"errors"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/mock"

	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
	inventoryv1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
)

func (s *APISuite) TestListPartsSuccess() {
	var (
		uuid    = gofakeit.UUID()
		filters = &inventoryv1.PartsFilter{Uuids: []string{uuid}}

		request = &inventoryv1.ListPartsRequest{Filter: filters}
		part    = &model.Part{
			Uuid: uuid,
		}
	)

	s.inventoryService.On("ListParts", s.ctx, mock.Anything).Return([]*model.Part{part}, nil)

	res, err := s.inventoryApi.ListParts(s.ctx, request)

	s.Require().NoError(err)
	s.Require().Len(res.Parts, 1)
	s.Require().True(res.Parts[0].Uuid == uuid)
}

func (s *APISuite) TestListPartsError() {
	var (
		uuid    = gofakeit.UUID()
		filters = &inventoryv1.PartsFilter{Uuids: []string{uuid}}

		request = &inventoryv1.ListPartsRequest{Filter: filters}
	)

	s.inventoryService.On("ListParts", s.ctx, mock.Anything).Return([]*model.Part{}, gofakeit.Error())

	res, err := s.inventoryApi.ListParts(s.ctx, request)

	s.Require().Error(err)
	s.Require().False(errors.Is(err, model.ErrPartNotFound))
	s.Require().NotNil(res)
}

func (s *APISuite) TestListPartsNotFound() {
	var (
		uuid    = gofakeit.UUID()
		filters = &inventoryv1.PartsFilter{Uuids: []string{uuid}}

		request = &inventoryv1.ListPartsRequest{Filter: filters}
	)

	s.inventoryService.On("ListParts", s.ctx, mock.Anything).Return([]*model.Part{}, model.ErrPartNotFound)

	res, err := s.inventoryApi.ListParts(s.ctx, request)

	s.Require().Error(err)
	s.Require().Nil(res)
}
