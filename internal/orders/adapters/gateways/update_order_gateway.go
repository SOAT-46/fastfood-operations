package gateways

import (
	"context"
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

func (itself *UpdateOrderGateway) Execute(ctx context.Context, order entities.Order) (*entities.Order, error) {
	dbOrder := models.BuildOrderFromDomain(order)
	updatedOrder, err := itself.repository.Update(ctx, dbOrder)
	if err != nil {
		return nil, fmt.Errorf("error updating order: %w", err)
	}

	return updatedOrder.ToDomain(), nil
}
