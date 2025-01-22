package gateways

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	global "github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
)

type GetOrdersPort interface {
	Execute(pagination global.Pagination) (global.PaginatedEntity[entities.Order], error)
}
