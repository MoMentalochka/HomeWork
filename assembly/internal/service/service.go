package service

import (
	"context"

	"github.com/MoMentalochka/HomeWork/assembly/internal/model"
)

type OrderPaidConsumerService interface {
	RunConsumer(ctx context.Context) error
}

type AssemblyProducerService interface {
	ProduceShipAssembled(ctx context.Context, event model.ShipAssembledEvent) error
}
