package contracts

import (
	"context"

	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
)

type CreateOrder interface {
	Execute(ctx context.Context, input entities.CreateOrderInput, listeners CreateOrderListeners)
}

type CreateOrderListeners struct {
	OnSuccess func(order entities.Order)
	OnInvalid func(err error)
	OnError   func(err error)
}
