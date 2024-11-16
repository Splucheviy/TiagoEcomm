package cart

import (
	"fmt"

	"github.com/Splucheviy/TiagoEcomm/types"
)

// getCartItemIDs ...
func getCartItemIDs(items []types.CartItem) ([]int, error) {
	productIDs := make([]int, 0, len(items))
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product: %d", item.ProductID)
		}

		productIDs[i] = item.ProductID
	}

	return productIDs, nil
}

// CreateOrder ...
func (h *Handler) CreateOrder(ps []types.Product, items []types.CartItem, userID int) (int, float64, error) {
	productMap := make(map[int]types.Product)

	for _, product := range ps {
		productMap[product.ID] = product
	}

	if err := checkIfCartIsInStock(items, productMap); err != nil {
		return 0, 0, nil
	}
	totalPrice := calculateTotalPrice(items, productMap)

	for _, item := range items {
		product, ok := productMap[item.ProductID]
		if !ok {
			continue
		}

		product.Quantity -= item.Quantity

		h.productStore.UpdateProduct(product)
	}

	orderID, err := h.store.CreateOrder(types.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "Unknown", // TODO: Implement address validation and storage in the database.
	})

	if err!= nil {
        return 0, 0, err
    }

	for _, item := range items {
		h.store.CreateOrderItem(types.OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productMap[item.ProductID].Price,
		})
	}
	
	return orderID, totalPrice, nil
}

// checkIfCartIsInStock ...
func checkIfCartIsInStock(cartItems []types.CartItem, products map[int]types.Product) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range cartItems {
		products, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("product not found: %d", item.ProductID)
		}

		if products.Quantity < item.Quantity {
			return fmt.Errorf("out of stock for product: %d", item.ProductID)
		}
	}

	return nil
}

// calculateTotalPrice ...
func calculateTotalPrice(cartItems []types.CartItem, products map[int]types.Product) float64 {
	totalPrice := 0.0

	for _, item := range cartItems {
		product, ok := products[item.ProductID]
		if !ok {
			continue
		}

		totalPrice += float64(item.Quantity) * product.Price
	}

	return totalPrice
}
