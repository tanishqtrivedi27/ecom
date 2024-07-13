package types

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(Id int) (*User, error)
	CreateUser(User) error
}

type ProductStore interface {
	GetProducts() ([]*Product, error)
	GetProductByID(productID int) (*Product, error)
	GetProductByIDs(productIDs []int) ([]Product, error)
	CreateProduct(CreateProductPayload) error
	UpdateProduct(Product) error
}

type OrderStore interface {
	CreateOrder(Order) (int, error)
	CreateOrderItem(OrderItem) error
	GetOrders(int) ([]*Order, error)
	GetOrderStatus(int, int) (string, error)
	CancelOrder(int, int) error
}
