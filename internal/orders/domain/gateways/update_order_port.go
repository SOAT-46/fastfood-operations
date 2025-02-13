package gateways

import (
	"context"

	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
)

type UpdateOrderPort interface {
	Execute(ctx context.Context, order entities.Order) (*entities.Order, error)
}
