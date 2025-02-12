package requests

import "github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"

type ProductRequest struct {
	Quantity int    `json:"quantity" binding:"required"`
	Product  string `json:"product" binding:"required"`
} // @name ProductRequest

func (itself ProductRequest) ToDomain() entities.OrderProduct {
	return entities.OrderProduct{
		Quantity: itself.Quantity,
		Product:  itself.Product,
	}
}
