package doubles

import (
	"context"
	"errors"

	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/contracts"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	"github.com/SOAT-46/fastfood-operations/test/orders/domain/builders"
)

var ErrUpdateOrderUseCaseStub = errors.New("test error")

type UpdateOrderUseCaseStub struct {
	callback func(listeners contracts.UpdateOrderListeners)
}

func NewUpdateOrderUseCaseStub() *UpdateOrderUseCaseStub {
	return &UpdateOrderUseCaseStub{}
}

func (itself *UpdateOrderUseCaseStub) WithOnError() *UpdateOrderUseCaseStub {
	itself.callback = func(listeners contracts.UpdateOrderListeners) {
		listeners.OnError(ErrUpdateOrderUseCaseStub)
	}
	return itself
}

func (itself *UpdateOrderUseCaseStub) WithOnNotFound() *UpdateOrderUseCaseStub {
	itself.callback = func(listeners contracts.UpdateOrderListeners) {
		listeners.OnNotFound()
	}
	return itself
}

func (itself *UpdateOrderUseCaseStub) WithOnSuccess() *UpdateOrderUseCaseStub {
	itself.callback = func(listeners contracts.UpdateOrderListeners) {
		order := builders.NewOrderBuilder().Build()
		listeners.OnSuccess(order)
	}
	return itself
}

func (itself *UpdateOrderUseCaseStub) Execute(
	_ context.Context, _ entities.Order, listeners contracts.UpdateOrderListeners) {
	itself.callback(listeners)
}
