package doubles

import (
	"errors"
	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/contracts"
	"github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
	"github.com/SOAT-46/fastfood-operations/test/orders/domain/builders"
)

var ErrGetOrdersUseCaseStub = errors.New("test error")

type GetOrdersUseCaseStub struct {
	callback func(listeners contracts.GetOrdersListeners)
}

func NewGetOrdersUseCaseStub() *GetOrdersUseCaseStub {
	return &GetOrdersUseCaseStub{}
}

func (itself *GetOrdersUseCaseStub) WithOnSuccess() *GetOrdersUseCaseStub {
	itself.callback = func(listeners contracts.GetOrdersListeners) {
		page := builders.NewOrderBuilder().BuildPaginated()
		listeners.OnSuccess(page)
	}
	return itself
}

func (itself *GetOrdersUseCaseStub) WithOnError() *GetOrdersUseCaseStub {
	itself.callback = func(listeners contracts.GetOrdersListeners) {
		listeners.OnError(ErrGetOrdersUseCaseStub)
	}
	return itself
}

func (itself *GetOrdersUseCaseStub) Execute(_ entities.Pagination, listeners contracts.GetOrdersListeners) {
	itself.callback(listeners)
}
