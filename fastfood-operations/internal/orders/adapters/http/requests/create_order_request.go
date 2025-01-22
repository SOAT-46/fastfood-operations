package requests

import "github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"

type CreateOrderRequest struct {
	UserID   *int             `json:"userId,omitempty"`
	Products []ProductRequest `json:"products" binding:"required"`
} // @name CreateOrderRequest

func (itself CreateOrderRequest) ToDomain() entities.CreateOrderInput {
	return entities.CreateOrderInput{
		Products: itself.buildProducts(),
		UserID:   itself.UserID,
	}
}

func (itself CreateOrderRequest) buildProducts() []entities.OrderProduct {
	var products []entities.OrderProduct
	for _, product := range itself.Products {
		products = append(products, product.ToDomain())
	}

	return products
}
