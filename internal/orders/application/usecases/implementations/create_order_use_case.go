package implementations

import (
	"context"
	"errors"

	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/contracts"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/gateways"
)

var ErrCreateOrderUseCase = errors.New("invalid order, impossible to save")

type CreateOrderUseCase struct {
	port gateways.SaveOrderPort
}

func NewCreateOrderUseCase(port gateways.SaveOrderPort) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		port: port,
	}
}

func (itself *CreateOrderUseCase) Execute(
	ctx context.Context, input entities.CreateOrderInput, listeners contracts.CreateOrderListeners) {
	if !input.IsValid() {
		listeners.OnInvalid(ErrCreateOrderUseCase)
		return
	}

	order, err := itself.port.Execute(ctx, input)
	if err != nil {
		listeners.OnError(err)
		return
	}
	listeners.OnSuccess(*order)
}
