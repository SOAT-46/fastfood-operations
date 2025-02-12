package doubles

import (
	"context"

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

func (itself *InMemoryGetOrderByIDPort) Execute(_ context.Context, _ string) (*entities.Order, error) {
	return itself.order, itself.err
}
