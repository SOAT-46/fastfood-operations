package contracts

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	global "github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
)

type GetOrders interface {
	Execute(pagination global.Pagination, listeners GetOrdersListeners)
}

type GetOrdersListeners struct {
	OnSuccess func(orders global.PaginatedEntity[entities.Order])
	OnError   func(err error)
}
