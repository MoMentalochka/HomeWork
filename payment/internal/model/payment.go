package model

type PayOrderRequest struct {
	OrderUuid     string
	UserUuid      string
	PaymentMethod string
}

type PayOrderResponse struct {
	TransactionUuid string
}
