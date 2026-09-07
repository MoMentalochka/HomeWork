package user

import "database/sql"

type orderRepository struct {
	db *sql.DB
}

func NewIamRepository(db *sql.DB) *orderRepository {
	return &orderRepository{
		db: db,
	}
}
