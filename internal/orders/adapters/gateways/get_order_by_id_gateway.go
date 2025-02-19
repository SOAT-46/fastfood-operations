package gateways

import (
	"context"
	"fmt"

	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/repositories/contracts"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
)

type GetOrderByIDGateway struct {
	repository contracts.OrdersRepository
}

func NewGetOrderByIDGateway(repository contracts.OrdersRepository) *GetOrderByIDGateway {
	return &GetOrderByIDGateway{repository: repository}
}

func (itself *GetOrderByIDGateway) Execute(ctx context.Context, id string) (*entities.Order, error) {
	order, err := itself.repository.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error getting order: %w", err)
	}
	return order.ToDomain(), nil
}
