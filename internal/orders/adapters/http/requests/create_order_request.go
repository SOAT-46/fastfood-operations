package requests

import "github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"

type CreateOrderRequest struct {
	UserID    *string          `json:"userId,omitempty"`
	PaymentID string           `json:"paymentId" binding:"required"`
	Products  []ProductRequest `json:"products" binding:"required"`
} // @name CreateOrderRequest

func (itself CreateOrderRequest) ToDomain() entities.CreateOrderInput {
	return entities.CreateOrderInput{
		Products:  itself.buildProducts(),
		UserID:    itself.UserID,
		PaymentID: itself.PaymentID,
	}
}

func (itself CreateOrderRequest) buildProducts() []entities.OrderProduct {
	var products []entities.OrderProduct
	for _, product := range itself.Products {
		products = append(products, product.ToDomain())
	}

	return products
}
