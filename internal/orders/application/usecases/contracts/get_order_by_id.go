package contracts

import (
	"context"

	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
)

type GetOrderByID interface {
	Execute(ctx context.Context, id string, listeners GetOrderByIDListeners)
}

type GetOrderByIDListeners struct {
	OnSuccess  func(order entities.Order)
	OnNotFound func()
	OnError    func(err error)
}
