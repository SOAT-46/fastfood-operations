package doubles

import (
	"errors"

	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
)

var ErrUpdateOrderPort = errors.New("error updating order")

type InMemoryUpdateOrderPort struct {
	order *entities.Order
	err   error
}

func NewInMemoryUpdateOrderPort() *InMemoryUpdateOrderPort {
	return &InMemoryUpdateOrderPort{}
}

func (itself *InMemoryUpdateOrderPort) WithOrder(order *entities.Order) *InMemoryUpdateOrderPort {
	itself.order = order
	return itself
}

func (itself *InMemoryUpdateOrderPort) WithError() *InMemoryUpdateOrderPort {
	itself.err = ErrUpdateOrderPort
	return itself
}

func (itself *InMemoryUpdateOrderPort) Execute(_ entities.Order) (*entities.Order, error) {
	return itself.order, itself.err
}
