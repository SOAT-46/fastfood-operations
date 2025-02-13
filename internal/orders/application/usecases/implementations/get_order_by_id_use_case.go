package implementations

import (
	"context"
	"errors"

	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/contracts"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/exceptions"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/gateways"
)

type GetOrderByIDUseCase struct {
	port gateways.GetOrderByIDPort
}

func NewGetOrderByIDUseCase(port gateways.GetOrderByIDPort) *GetOrderByIDUseCase {
	return &GetOrderByIDUseCase{
		port: port,
	}
}

func (itself *GetOrderByIDUseCase) Execute(
	ctx context.Context, id string, listeners contracts.GetOrderByIDListeners) {
	order, err := itself.port.Execute(ctx, id)
	if err != nil {
		if errors.Is(err, exceptions.ErrOrderNotFound) {
			listeners.OnNotFound()
		} else {
			listeners.OnError(err)
		}
		return
	}

	listeners.OnSuccess(*order)
}
