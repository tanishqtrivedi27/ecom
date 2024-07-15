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

	err := s.db.QueryRow("INSERT INTO orders (userId, billingAddressId, total, status) VALUES ($1, $2, $3, $4) RETURNING id", order.UserID, order.BillingAddressID, order.Total, order.Status).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}

func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	_, err := s.db.Exec("INSERT INTO order_items (orderId, productId, quantity, price) VALUES ($1, $2, $3, $4)", orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
	return err
}

func (s *Store) GetOrders(userId int) ([]*types.Order, error) {
	rows, err := s.db.Query("SELECT * FROM orders WHERE userId = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]*types.Order, 0)
	for rows.Next() {
		p, err := scanRowsIntoOrder(rows)
		if err != nil {
			return nil, err
		}

		orders = append(orders, p)
	}

	return orders, nil
}

func (s *Store) GetOrderStatus(userId, orderId int) (string, error) {
	var status string
	row := s.db.QueryRow("SELECT status FROM orders WHERE userId = $1 AND id = $2", userId, orderId)
	err := row.Scan(&status)

	if err != nil {
		return "", err
	}

	return status, nil
}


func (s *Store) CancelOrder(userId, orderId int) error {
	
	_, err := s.db.Exec("UPDATE orders SET status = 'cancelled' WHERE userId = $1 AND id = $2", userId, orderId)
	return err
}

func scanRowsIntoOrder(rows *sql.Rows) (*types.Order, error) {
	order := new(types.Order)

	err := rows.Scan(
		&order.ID,
		&order.UserID,
		&order.BillingAddressID,
		&order.Total,
		&order.Status,
		&order.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return order, nil
}
