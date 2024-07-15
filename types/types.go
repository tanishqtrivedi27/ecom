package types

type RegisterUserPayLoad struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}

type LoginUserPayLoad struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CartCheckoutItem struct {
	ProductID int `json:"productID"`
	Quantity  int `json:"quantity"`
}

type CartCheckoutPayload struct {
	Items            []CartCheckoutItem `json:"items" validate:"required"`
	BillingAddressID int                `json:"billingAddressId" validate:"required"`
}

type CreateProductPayload struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
}

type CreateAddressPayload struct {
	Line1   string `json:"line1" validate:"required"`
	Line2   string `json:"line2"`
	City    string `json:"city" validate:"required"`
	Country string `json:"country" validate:"required"`
}
