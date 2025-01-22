package doubles

import (
	"errors"

	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/contracts"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	"github.com/SOAT-46/fastfood-operations/test/orders/domain/builders"
)

var ErrCreateOrderUseCaseStub = errors.New("test error")

type CreateOrderUseCaseStub struct {
	callback func(listeners contracts.CreateOrderListeners)
}

func NewCreateOrderUseCaseStub() *CreateOrderUseCaseStub {
	return &CreateOrderUseCaseStub{}
}

func (itself *CreateOrderUseCaseStub) WithOnSuccess() *CreateOrderUseCaseStub {
	itself.callback = func(listeners contracts.CreateOrderListeners) {
		order := builders.NewOrderBuilder().Build()
		listeners.OnSuccess(order)
	}
	return itself
}

func (itself *CreateOrderUseCaseStub) WithOnInvalid() *CreateOrderUseCaseStub {
	itself.callback = func(listeners contracts.CreateOrderListeners) {
		listeners.OnInvalid(ErrCreateOrderUseCaseStub)
	}
	return itself
}

func (itself *CreateOrderUseCaseStub) WithOnError() *CreateOrderUseCaseStub {
	itself.callback = func(listeners contracts.CreateOrderListeners) {
		listeners.OnError(ErrCreateOrderUseCaseStub)
	}
	return itself
}

func (itself *CreateOrderUseCaseStub) Execute(_ entities.CreateOrderInput, listeners contracts.CreateOrderListeners) {
	itself.callback(listeners)
}
