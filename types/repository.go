package types

import "database/sql"

type UserStore interface {
	GetUserByEmail(string) (*User, error)
	GetUserById(int) (*User, error)
	CreateUser(User) error
	UpdateLastLogin(int) error
	GetAddresses(int) ([]*Address, error)
	CreateAddress(Address) error
	CheckIfValidAddress(int, int) (bool, error)
}

type ProductStore interface {
	GetProducts() ([]*Product, error)
	GetProductByID(int) (*Product, error)
	CreateProduct(CreateProductPayload) error
	
	BeginTx() (*sql.Tx, error)
	GetProductByIDsTx(*sql.Tx, []int) ([]*Product, error)
	UpdateProductTx(*sql.Tx, *Product) error
}

type OrderStore interface {
	CreateOrder(Order) (int, error)
	CreateOrderItem(OrderItem) error
	GetOrders(int) ([]*Order, error)
	GetOrderStatus(int, int) (string, error)
	CancelOrder(int, int) error
}
