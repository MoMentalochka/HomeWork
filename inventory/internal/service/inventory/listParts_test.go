package inventory

import (
	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-faster/errors"
	"github.com/stretchr/testify/mock"
)

func (s *ServiceSuite) TestGetPartsNotFound() {

	var (
		uuid    = gofakeit.UUID()
		filters = &model.PartsFilter{Uuids: []string{uuid}}
	)

	s.inventoryRepository.On("ListParts", s.ctx, mock.Anything).Return([]*model.Part{}, model.ErrPartNotFound)

	_, err := s.inventoryService.ListParts(s.ctx, filters)

	s.Require().Error(err)
	s.Require().True(errors.Is(err, model.ErrPartNotFound))
}

func (s *ServiceSuite) TestGetPartsSuccess() {

	var (
		uuid    = gofakeit.UUID()
		filters = &model.PartsFilter{Uuids: []string{uuid}}
		part    = &model.Part{
			Uuid: uuid,
		}
	)

	s.inventoryRepository.On("ListParts", s.ctx, mock.Anything).Return([]*model.Part{part}, nil)

	res, err := s.inventoryService.ListParts(s.ctx, filters)

	s.Require().NoError(err)
	s.Require().Len(res, 1)
	s.Require().True(res[0].Uuid == uuid)

}
