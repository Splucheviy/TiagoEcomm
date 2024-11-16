package cart

import (
	"fmt"
	"net/http"

	"github.com/Splucheviy/TiagoEcomm/types"
	"github.com/Splucheviy/TiagoEcomm/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// Handler ...
type Handler struct {
	store        types.OrderStore
	productStore types.ProductStore
}

// NewHandler ...
func NewHandler(store types.OrderStore, productStore types.ProductStore) *Handler {
	return &Handler{store: store, productStore: productStore}
}

// RegisterRoutes ...
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", h.HandleCheckout).Methods(http.MethodPost)
}

// HandleCheckout ...
func (h *Handler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	userID := 0
	var cart types.CartCheckoutPayload
	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	productIDs, err := getCartItemIDs(cart.Items)
	if err!= nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

	ps, err := h.productStore.GetProductsByIDs(productIDs)
	if err!= nil {
        utils.WriteError(w, http.StatusInternalServerError, err)
        return
    }

	orderID, totalPrice, err := h.CreateOrder(ps, cart.Items, userID)
	if err!= nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
        "order_id": orderID,
        "total_price": totalPrice,
    })
}
