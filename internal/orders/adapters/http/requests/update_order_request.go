package requests

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
)

type UpdateOrderRequest struct {
	Status string `json:"status" binding:"required"`
} // @name UpdateOrderRequest

func (itself UpdateOrderRequest) ToDomain(id int) entities.Order {
	status := entities.ToOrderStatus(itself.Status)
	return entities.Order{
		ID:     id,
		Status: status,
	}
}
