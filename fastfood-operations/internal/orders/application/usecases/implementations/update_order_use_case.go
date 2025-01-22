package implementations

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/contracts"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/gateways"
)

type UpdateOrderUseCase struct {
	port    gateways.UpdateOrderPort
	getPort gateways.GetOrderByIDPort
}

func NewUpdateOrderUseCase(
	port gateways.UpdateOrderPort,
	getPort gateways.GetOrderByIDPort) *UpdateOrderUseCase {
	return &UpdateOrderUseCase{
		port:    port,
		getPort: getPort,
	}
}

func (itself *UpdateOrderUseCase) Execute(order entities.Order, listeners contracts.UpdateOrderListeners) {
	found, errFound := itself.getPort.Execute(order.ID)
	if errFound != nil {
		listeners.OnError(errFound)
		return
	}
	if found == nil {
		listeners.OnNotFound()
		return
	}

	updatedOrder, err := itself.port.Execute(order)
	if err != nil {
		listeners.OnError(err)
		return
	}
	listeners.OnSuccess(*updatedOrder)
}
