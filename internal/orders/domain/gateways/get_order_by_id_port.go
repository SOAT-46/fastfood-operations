package gateways

import "github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"

type GetOrderByIDPort interface {
	Execute(id int) (*entities.Order, error)
}
