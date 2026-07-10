package payment

import (
	"context"
	"testing"

	"github.com/MoMentalochka/HomeWork/payment/internal/model"
	"github.com/MoMentalochka/HomeWork/payment/internal/service/mocks"
	payment_v1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/payment/v1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestPayOrder(t *testing.T) {

	ctx := context.Background()
	id := uuid.New().String()

	client := mocks.NewPaymentService(t)
	client.On("PayOrder", ctx, &model.PayOrderRequest{}).Return(&payment_v1.PayOrderResponse{TransactionUuid: id}, nil)

	t.Run("Check valid uuid", func(t *testing.T) {
		res, err := client.PayOrder(ctx, &model.PayOrderRequest{})
		require.NoError(t, uuid.Validate(res.TransactionUuid))
		require.NoError(t, err)
	})

	t.Run("Check no error", func(t *testing.T) {
		_, err := client.PayOrder(ctx, &model.PayOrderRequest{})
		require.NoError(t, err)
	})

}
