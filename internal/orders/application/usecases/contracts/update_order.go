package contracts

import "github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"

type UpdateOrder interface {
	Execute(order entities.Order, listeners UpdateOrderListeners)
}

type UpdateOrderListeners struct {
	OnSuccess  func(order entities.Order)
	OnNotFound func()
	OnError    func(err error)
}
