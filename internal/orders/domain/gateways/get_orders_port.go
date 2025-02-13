package gateways

import (
	"context"

	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	global "github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
)

type GetOrdersPort interface {
	Execute(ctx context.Context, pagination global.Pagination) (global.PaginatedEntity[entities.Order], error)
}
