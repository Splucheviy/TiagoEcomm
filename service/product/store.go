package product

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Splucheviy/TiagoEcomm/types"
)

// Store ...
type Store struct {
	db *sql.DB
}

// NewStore ...
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// GetProducts ...
func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}

	return products, nil
}

// scanRowsIntoProduct ...
func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// GetProductsByIDs ...
func (s *Store) GetProductsByIDs(productIDs []int) ([]types.Product, error) {
	placeholders := strings.Repeat(",?", len(productIDs)-1)
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (?%s)", placeholders)

	args := make([]interface{}, 0, len(productIDs))
	for i, v := range productIDs {
		args[i] = v
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	products := []types.Product{}
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}

	return products, nil

}

// UpdateProduct ...
func (s *Store) UpdateProduct(product types.Product) error {
	query := "UPDATE products SET name=?, description=?, image=?, price=?, quantity=? WHERE id=?"
	_, err := s.db.Exec(query,
		product.Name,
		product.Description,
		product.Image,
		product.Price,
		product.Quantity,
		product.ID,
	)
	if err != nil {
		return err
	}

	return nil
}
