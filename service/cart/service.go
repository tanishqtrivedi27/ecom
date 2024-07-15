package cart

import (
	"fmt"
	"sync/atomic"

	"github.com/tanishqtrivedi27/ecom/types"
)

func getCartItemIds(items []types.CartCheckoutItem) ([]int, error) {
	productIds := make([]int, len(items))

	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product %v", item.ProductID)
		}

		productIds[i] = item.ProductID
	}

	return productIds, nil
}

func (h *Handler) createOrder(ps []*types.Product, items []types.CartCheckoutItem, userId, billingAddressId int) (int, float64, error) {
	productMap := make(map[int]*types.Product)
	for _, product := range ps {
		productMap[product.ID] = product
	}

	// check availability
	if err := checkIfInStock(items, productMap); err != nil {
		return 0, 0, err
	}

	// calc total price
	totalPrice := calculateTotalPrice(items, productMap)

	// update quantity in products table
	for _, item := range items {
		product := productMap[item.ProductID]
		atomic.AddInt32(&product.Quantity, -int32(item.Quantity))
		h.productStore.UpdateProduct(*product)
	}

	//create records in order, order_items table
	orderID, err := h.orderStore.CreateOrder(types.Order{
		UserID:           userId,
		BillingAddressID: billingAddressId,
		Total:            totalPrice,
		Status:           "pending",
	})
	
	if err != nil {
		return 0, 0, err
	}

	for _, item := range items {
		h.orderStore.CreateOrderItem(types.OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productMap[item.ProductID].Price,
		})
	}

	return orderID, totalPrice, nil
}

func checkIfInStock(cartItems []types.CartCheckoutItem, products map[int]*types.Product) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range cartItems {
		product, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("product %v is not available in the store, please refresh your cart", item.ProductID)
		}

		if atomic.LoadInt32(&product.Quantity) < int32(item.Quantity) {
			return fmt.Errorf("product %v is not available in the quantity requested", product.Name)
		}
	}

	return nil
}

func calculateTotalPrice(cartItems []types.CartCheckoutItem, products map[int]*types.Product) float64 {
	var total float64

	for _, item := range cartItems {
		product := products[item.ProductID]
		total += product.Price * float64(item.Quantity)
	}

	return total
}
