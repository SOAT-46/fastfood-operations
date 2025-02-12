package gateways

import (
	"context"

	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
)

type GetOrderByIDPort interface {
	Execute(ctx context.Context, id string) (*entities.Order, error)
}
