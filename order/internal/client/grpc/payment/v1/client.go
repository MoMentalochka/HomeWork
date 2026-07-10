package v1

import (
	"context"

	paymentv1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/payment/v1"
)

type paymentClient struct {
	generatedClient paymentv1.PaymentServiceClient
}

func NewPaymentClient(generatedClient paymentv1.PaymentServiceClient) *paymentClient {
	return &paymentClient{generatedClient: generatedClient}
}

func (p *paymentClient) PayOrder(ctx context.Context, req *paymentv1.PayOrderRequest) (string, error) {
	res, err := p.generatedClient.PayOrder(ctx, req)
	if err != nil {
		return "", err
	}
	return res.TransactionUuid, nil
}
