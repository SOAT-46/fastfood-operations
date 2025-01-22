package gateways

import (
	"fmt"
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/repositories/contracts"
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/repositories/models"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
)

type UpdateOrderGateway struct {
	repository contracts.OrdersRepository
}

func NewUpdateOrderGateway(repository contracts.OrdersRepository) *UpdateOrderGateway {
	return &UpdateOrderGateway{
		repository: repository,
	}
}

func (itself *UpdateOrderGateway) Execute(order entities.Order) (*entities.Order, error) {
	dbOrder := models.BuildGormOrderFromDomain(order)
	updatedOrder, err := itself.repository.Update(dbOrder)
	if err != nil {
		return nil, fmt.Errorf("error updating order: %w", err)
	}
	return updatedOrder.ToDomain(), nil
}
