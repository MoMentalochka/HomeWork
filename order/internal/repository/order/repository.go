package order

import (
	"database/sql"
	"log"
	"strings"

	"github.com/MoMentalochka/HomeWork/order/internal/model"
)

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *orderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) Create(order *model.OrderDto) error {
	_, err := r.db.Exec(
		"INSERT INTO orders (order_uuid, user_uuid, part_uuids, status) VALUES ($1, $2, $3, $4)",
		order.OrderUUID,
		order.UserUUID,
		strings.Join(order.PartUuids, ","),
		order.Status,
	)
	if err != nil {
		log.Printf("failed to insert note: %v\n", err)
		return err
	}
	return nil
}

func (r *orderRepository) Get(id string) (*model.OrderDto, error) {

	rows, err := r.db.Query("SELECT * FROM orders WHERE order_uuid = $1", id)
	if err != nil {
		log.Printf("failed to select order: %v\n", err)
		return &model.OrderDto{}, err
	}
	defer func() {
		err = rows.Close()
		log.Printf("rows closed: %v\n", err)
	}()
	var idx int
	var transaction_uuid, payment_method, part_uuids sql.NullString
	var OrderUUID, UserUUID, Status string
	var total_price sql.NullFloat64

	var PartUuids []string
	var TransactionUUID, PaymentMethod string
	var TotalPrice float64
	for rows.Next() {
		err = rows.Scan(
			&idx,
			&OrderUUID,
			&UserUUID,
			&part_uuids,
			&total_price,
			&transaction_uuid,
			&payment_method,
			&Status,
		)
		if err != nil {
			return &model.OrderDto{}, err
		}
	}
	if OrderUUID == "" {
		return &model.OrderDto{}, model.ErrPartNotFound
	}
	// Провека потенциально пустых значений
	if part_uuids.Valid {
		PartUuids = strings.Split(part_uuids.String, ",")
	}
	if transaction_uuid.Valid {
		TransactionUUID = transaction_uuid.String
	}

	if payment_method.Valid {
		PaymentMethod = payment_method.String
	}
	if total_price.Valid {
		TotalPrice = total_price.Float64
	}

	return &model.OrderDto{
		OrderUUID:       OrderUUID,
		UserUUID:        UserUUID,
		PartUuids:       PartUuids,
		TotalPrice:      TotalPrice,
		TransactionUUID: TransactionUUID,
		PaymentMethod:   PaymentMethod,
		Status:          Status,
	}, nil
}

func (r *orderRepository) Update(id string, data model.OrderDto) (*model.OrderDto, error) {
	return nil, nil
}
