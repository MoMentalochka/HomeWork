package model

type OrderPaidEvent struct {
	OrderUUID       string
	EventUUID       string
	UserUUID        string
	PaymentMethod   string
	TransactionUUID string
}
