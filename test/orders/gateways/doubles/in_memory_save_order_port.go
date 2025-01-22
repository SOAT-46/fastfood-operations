package doubles

import (
	"errors"

	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
)

var ErrSaveOrder = errors.New("error saving order")

type InMemorySaveOrderPort struct {
	err   error
	order *entities.Order
}

func NewInMemorySaveOrderPort() *InMemorySaveOrderPort {
	return &InMemorySaveOrderPort{}
}

func (itself *InMemorySaveOrderPort) WithError() *InMemorySaveOrderPort {
	itself.err = ErrSaveOrder
	return itself
}

func (itself *InMemorySaveOrderPort) WithOrder(order *entities.Order) *InMemorySaveOrderPort {
	itself.order = order
	return itself
}

func (itself *InMemorySaveOrderPort) Execute(_ entities.CreateOrderInput) (*entities.Order, error) {
	return itself.order, itself.err
}
