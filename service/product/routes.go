package product

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/tanishqtrivedi27/ecom/types"
	"github.com/tanishqtrivedi27/ecom/utils"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /products", h.handleGetProducts)
	router.HandleFunc("GET /products/{productID}", h.handleGetProduct)
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, ps)
}

func (h *Handler) handleGetProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := strconv.Atoi(r.PathValue("productID"))

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid product ID"))
		return
	}

	product, err := h.store.GetProductByID(productID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, product)
}
