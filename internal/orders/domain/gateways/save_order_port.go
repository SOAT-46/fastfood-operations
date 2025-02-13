package gateways

import (
	"context"

	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
)

type SaveOrderPort interface {
	Execute(ctx context.Context, order entities.CreateOrderInput) (*entities.Order, error)
}
