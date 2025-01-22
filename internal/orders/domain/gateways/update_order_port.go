package gateways

import "github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"

type UpdateOrderPort interface {
	Execute(order entities.Order) (*entities.Order, error)
}
