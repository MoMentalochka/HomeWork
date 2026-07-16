package order

import (
	"database/sql"
	"log"

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

func (r *orderRepository) AddOrder(order *model.OrderDto) error {
	_, err := r.db.Exec(
		"INSERT INTO orders (order_uuid, user_uuid, part_uuids, status) VALUES ($1, $2, $3, $4)",
		order.OrderUUID,
		order.UserUUID,
		order.PartUuids,
		order.Status,
	)
	if err != nil {
		log.Printf("failed to insert note: %v\n", err)
		return err
	}
	return nil
}

func (r *orderRepository) GetOrder(id string) (*model.OrderDto, error) {

	var response model.OrderDto
	rows, err := r.db.Query("SELECT * FROM orders WHERE order_uuid = $1", id)
	if err != nil {
		log.Printf("failed to insert note: %v\n", err)
		return &model.OrderDto{}, err
	}
	defer func() {
		err = rows.Close()
		log.Printf("rows closed: %v\n", err)
	}()

	err = rows.Scan(
		&response.OrderUUID,
		&response.UserUUID,
		&response.PartUuids,
		&response.TotalPrice,
		&response.TransactionUUID,
		&response.PaymentMethod,
		&response.Status,
	)
	if err != nil {
		return &model.OrderDto{}, err
	}

	return &response, nil
}
