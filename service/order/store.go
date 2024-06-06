package order

import (
	"database/sql"

	"github.com/tanishqtrivedi27/ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(order types.Order) (int, error) {
	var lastInsertId int

	err := s.db.QueryRow("INSERT INTO orders (userId, total, status, address) VALUES ($1, $2, $3, $4) RETURNING id", order.UserID, order.Total, order.Status, order.Address).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}


func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	_, err := s.db.Exec("INSERT INTO order_items (orderId, productId, quantity, price) VALUES ($1, $2, $3, $4)", orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
	return err
}