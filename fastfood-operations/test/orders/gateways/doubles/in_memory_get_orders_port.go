package doubles

import (
	"errors"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	global "github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
)

var ErrNoOrders = errors.New("no orders")

type InMemoryGetOrdersPort struct {
	page global.PaginatedEntity[entities.Order]
	err  error
}

func NewInMemoryGetOrdersPort() *InMemoryGetOrdersPort {
	return &InMemoryGetOrdersPort{}
}

func (itself *InMemoryGetOrdersPort) WithOrders(orders global.PaginatedEntity[entities.Order]) *InMemoryGetOrdersPort {
	itself.page = orders
	return itself
}

func (itself *InMemoryGetOrdersPort) WithError() *InMemoryGetOrdersPort {
	itself.err = ErrNoOrders
	return itself
}

func (itself *InMemoryGetOrdersPort) Execute(_ global.Pagination) (global.PaginatedEntity[entities.Order], error) {
	return itself.page, itself.err
}
