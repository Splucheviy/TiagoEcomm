package types

import "time"

// ProductStore ...
type ProductStore interface {
	GetProducts() ([]Product, error)
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
