package inventory

import (
	"errors"

	"github.com/MoMentalochka/HomeWork/inventory/internal/model"
	repomodel "github.com/MoMentalochka/HomeWork/inventory/internal/repository/model"
	"github.com/brianvoe/gofakeit/v7"
)

func (s *ServiceSuite) TestGetPartSuccess() {

	var (
		uuid = gofakeit.UUID()
		name = gofakeit.Name()

		repoPart = repomodel.Part{
			Uuid: uuid,
			Name: name,
		}
	)

	s.inventoryRepository.On("GetPart", s.ctx, uuid).Return(repoPart, nil)

	part, err := s.inventoryService.GetPart(s.ctx, uuid)

	s.Require().NoError(err)
	s.Require().Equal(uuid, part.Uuid)
	s.Require().Equal(name, part.Name)

}

func (s *ServiceSuite) TestGetPartNotFound() {

	var (
		uuid     = gofakeit.UUID()
		repoPart = repomodel.Part{}
	)

	s.inventoryRepository.On("GetPart", s.ctx, uuid).Return(repoPart, model.ErrPartNotFound)

	part, err := s.inventoryService.GetPart(s.ctx, uuid)

	s.Require().Error(err)
	s.Require().True(errors.Is(err, model.ErrPartNotFound))
	s.Require().Empty(part)
}

func (s *ServiceSuite) TestGetPartError() {

	var (
		uuid     = gofakeit.UUID()
		repoPart = repomodel.Part{}
	)

	s.inventoryRepository.On("GetPart", s.ctx, uuid).Return(repoPart, gofakeit.Error())

	part, err := s.inventoryService.GetPart(s.ctx, uuid)

	s.Require().Error(err)
	s.Require().False(errors.Is(err, model.ErrPartNotFound))
	s.Require().Empty(part)
}
