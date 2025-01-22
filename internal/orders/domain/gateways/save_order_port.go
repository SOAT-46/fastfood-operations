package gateways

import "github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"

type SaveOrderPort interface {
	Execute(order entities.CreateOrderInput) (*entities.Order, error)
}
