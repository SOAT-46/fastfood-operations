package gateways

import (
	"context"
	"fmt"

	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/repositories/contracts"
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/repositories/models"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
)

type SaveOrderGateway struct {
	repository contracts.OrdersRepository
}

func NewSaveOrderGateway(repository contracts.OrdersRepository) *SaveOrderGateway {
	return &SaveOrderGateway{repository: repository}
}

func (itself *SaveOrderGateway) Execute(
	ctx context.Context, order entities.CreateOrderInput) (*entities.Order, error) {
	dbOrder := models.BuildOrderFromDomainInput(order)
	newOrder, err := itself.repository.Save(ctx, dbOrder)
	if err != nil {
		return nil, fmt.Errorf("error saving order: %w", err)
	}
	return newOrder.ToDomain(), nil
}
