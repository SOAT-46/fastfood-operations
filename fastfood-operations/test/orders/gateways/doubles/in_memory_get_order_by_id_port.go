package doubles

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
)

type InMemoryGetOrderByIDPort struct {
	order *entities.Order
	err   error
}

func NewInMemoryGetOrderByIDPort() *InMemoryGetOrderByIDPort {
	return &InMemoryGetOrderByIDPort{}
}

func (itself *InMemoryGetOrderByIDPort) WithError(err error) *InMemoryGetOrderByIDPort {
	itself.err = err
	return itself
}

func (itself *InMemoryGetOrderByIDPort) WithOrder(order *entities.Order) *InMemoryGetOrderByIDPort {
	itself.order = order
	return itself
}

func (itself *InMemoryGetOrderByIDPort) Execute(_ int) (*entities.Order, error) {
	return itself.order, itself.err
}
