package types

import "time"

// ProductStore ...
type ProductStore interface {
	GetProducts() ([]Product, error)
	GetProductsByIDs(ps []int) ([]Product, error)
	UpdateProduct(Product) error
}

// OrderStore ...
type OrderStore interface {
	CreateOrder(Order) (int, error)
	CreateOrderItem(OrderItem) error
}

// Order ...
type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userId"`
	Total     float64   `json:"total"`
	Status    string    `json:"status"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
}

// OrderItem ...
type OrderItem struct {
	ID        int       `json:"id"`
	OrderID   int       `json:"orderId"`
	ProductID int       `json:"productId"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}

// Product ...
type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"createdAt"`
}

// LoginUserPayload struct ...
type LoginUserPayload struct {
	Email    string `json:"email"  validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// RegisterUserPayload struct ...
type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email"  validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=255"`
}

// User struct ...
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

// UserStore interface
type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

// CartItem ...
type CartItem struct {
	ProductID int `json:"productID"`
	Quantity  int `json:"quantity"`
}

// CartCheckoutPayload ...
type CartCheckoutPayload struct {
	Items []CartItem `json:"items" validate:"required"`
}
