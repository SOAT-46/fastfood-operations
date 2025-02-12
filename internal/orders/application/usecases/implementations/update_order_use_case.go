package implementations

import (
	"context"

	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/contracts"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/gateways"
)

type UpdateOrderUseCase struct {
	port            gateways.UpdateOrderPort
	getOrderUseCase contracts.GetOrderByID
}

func NewUpdateOrderUseCase(
	port gateways.UpdateOrderPort,
	getOrderUseCase contracts.GetOrderByID) *UpdateOrderUseCase {
	return &UpdateOrderUseCase{
		port:            port,
		getOrderUseCase: getOrderUseCase,
	}
}

func (itself *UpdateOrderUseCase) Execute(
	ctx context.Context, order entities.Order, listeners contracts.UpdateOrderListeners) {
	getOrderListeners := contracts.GetOrderByIDListeners{
		OnSuccess: func(found entities.Order) {
			itself.onFound(ctx, found, order.Status, listeners)
		},
		OnNotFound: listeners.OnNotFound,
		OnError:    listeners.OnError,
	}
	itself.getOrderUseCase.Execute(ctx, order.Number, getOrderListeners)
}

func (itself *UpdateOrderUseCase) onFound(
	ctx context.Context,
	found entities.Order,
	newStatus entities.OrderStatus,
	listeners contracts.UpdateOrderListeners) {
	found.Status = newStatus
	updatedOrder, err := itself.port.Execute(ctx, found)
	if err != nil {
		listeners.OnError(err)
		return
	}
	listeners.OnSuccess(*updatedOrder)
}
