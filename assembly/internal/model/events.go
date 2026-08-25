package model

type OrderPaidEvent struct {
	OrderUUID       string
	EventUUID       string
	UserUUID        string
	PaymentMethod   string
	TransactionUUID string
}

type ShipAssembledEvent struct {
	EventUUID    string
	OrderUUID    string
	UserUUID     string
	BuildTimeSec int64
}
