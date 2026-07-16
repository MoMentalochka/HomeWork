package model

type OrderDto struct {
	OrderUUID string
	// User ID.
	UserUUID string
	// Parts Id array.
	PartUuids []string
	// Total price.
	TotalPrice float64
	// User ID.
	TransactionUUID string
	PaymentMethod   string
	Status          string
}
