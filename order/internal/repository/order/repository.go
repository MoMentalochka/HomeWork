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
		"INSERT INTO orders (order_uuid, user_uuid, part_uuids, total_price, status) VALUES ($1, $2, $3, $4, $5)",
		order.OrderUUID,
		order.UserUUID,
		strings.Join(order.PartUuids, ","),
		order.TotalPrice,
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
		if err != nil {
			log.Printf("fail to close rows: %v\n", err)
		}
	}()

	// Извлекаем данные в модель
	{
		var idx int
		var transactionUuid, paymentMethod, partUuids sql.NullString
		var OrderUUID, UserUUID, Status string
		var totalPrice sql.NullFloat64

		var PartUuids []string
		var TransactionUUID, PaymentMethod string
		var TotalPrice float64
		for rows.Next() {
			err = rows.Scan(
				&idx,
				&OrderUUID,
				&UserUUID,
				&partUuids,
				&totalPrice,
				&transactionUuid,
				&paymentMethod,
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
		if partUuids.Valid {
			PartUuids = strings.Split(partUuids.String, ",")
		}
		if transactionUuid.Valid {
			TransactionUUID = transactionUuid.String
		}

		if paymentMethod.Valid {
			PaymentMethod = paymentMethod.String
		}
		if totalPrice.Valid {
			TotalPrice = totalPrice.Float64
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
}

func (r *orderRepository) Update(id string, data *model.OrderDto) error {
	_, err := r.db.Exec("UPDATE orders SET status = $1, payment_method = NULLIF($2, ''), transaction_uuid = NULLIF($3, '') WHERE order_uuid = $4", data.Status, data.PaymentMethod, data.TransactionUUID, id)
	if err != nil {
		log.Printf("failed to update order: %v\n", err)
		return err
	}
	return nil
}
