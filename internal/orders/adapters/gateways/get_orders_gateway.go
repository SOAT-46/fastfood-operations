package gateways

import (
	"context"

	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/gateways/mappers"
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/repositories/contracts"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	global "github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
)

type GetOrdersGateway struct {
	repository contracts.OrdersRepository
}

func NewGetOrdersGateway(repository contracts.OrdersRepository) *GetOrdersGateway {
	return &GetOrdersGateway{
		repository: repository,
	}
}

func (itself *GetOrdersGateway) Execute(
	ctx context.Context, pagination global.Pagination) (global.PaginatedEntity[entities.Order], error) {
	orders, err := itself.repository.ListAll(ctx, pagination)
	if err != nil {
		return global.NewPaginatedEntity([]entities.Order{}, pagination), err
	}
	return mappers.MapToPaginatedEntity(orders), nil
}
