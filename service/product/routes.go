package product

import (
	"net/http"

	"github.com/Splucheviy/TiagoEcomm/types"
	"github.com/Splucheviy/TiagoEcomm/utils"
	"github.com/gorilla/mux"
)

// Handler ...
type Handler struct {
	store types.ProductStore
}

// NewHandler ...
func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

// RegisterRoutes ...
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodGet)
}

// handleCreateProduct ..
func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, ps)
}
