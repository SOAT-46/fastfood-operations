package doubles

import (
	"context"

	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/contracts"
	"github.com/SOAT-46/fastfood-operations/test/orders/domain/builders"
)

type GetOrderByIDUseCaseStub struct {
	callback func(listeners contracts.GetOrderByIDListeners)
}

func NewGetOrderByIDUseCaseStub() *GetOrderByIDUseCaseStub {
	return &GetOrderByIDUseCaseStub{}
}

func (itself *GetOrderByIDUseCaseStub) WithOnSuccess() *GetOrderByIDUseCaseStub {
	itself.callback = func(listeners contracts.GetOrderByIDListeners) {
		order := builders.NewOrderBuilder().Build()
		listeners.OnSuccess(order)
	}
	return itself
}

func (itself *GetOrderByIDUseCaseStub) WithOnNotFound() *GetOrderByIDUseCaseStub {
	itself.callback = func(listeners contracts.GetOrderByIDListeners) {
		listeners.OnNotFound()
	}
	return itself
}

func (itself *GetOrderByIDUseCaseStub) WithOnError(err error) *GetOrderByIDUseCaseStub {
	itself.callback = func(listeners contracts.GetOrderByIDListeners) {
		listeners.OnError(err)
	}
	return itself
}

func (itself *GetOrderByIDUseCaseStub) Execute(
	_ context.Context, _ string, listeners contracts.GetOrderByIDListeners) {
	itself.callback(listeners)
}
