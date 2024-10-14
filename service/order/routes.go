package order

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/tanishqtrivedi27/ecom/service/auth"
	"github.com/tanishqtrivedi27/ecom/types"
	"github.com/tanishqtrivedi27/ecom/utils"
)

type Handler struct {
	orderStore types.OrderStore
	userStore  types.UserStore
}

func NewHandler(orderStore types.OrderStore, userStore types.UserStore) *Handler {
	return &Handler{orderStore: orderStore, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /orders", auth.JWTAuthMiddleware(http.HandlerFunc(h.getOrders)))
	router.HandleFunc("POST /orders/{id}/cancel", auth.JWTAuthMiddleware(http.HandlerFunc(h.cancelOrder)))
}

func (h *Handler) getOrders(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserFromContext(r.Context())

	orders, err := h.orderStore.GetOrders(userId)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, orders)
}

func (h *Handler) cancelOrder(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserFromContext(r.Context())
	orderId, _ := strconv.Atoi(r.PathValue("id"))

	orderStatus, err := h.orderStore.GetOrderStatus(userId, orderId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid order id"))
		return
	}

	if orderStatus != "pending" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("order has already been %v", orderStatus))
		return
	}

	err = h.orderStore.CancelOrder(userId, orderId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "order cancelled succesfully")
}
