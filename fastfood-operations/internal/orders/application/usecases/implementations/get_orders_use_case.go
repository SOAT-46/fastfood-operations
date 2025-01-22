package implementations

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/contracts"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/gateways"
	global "github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
)

type GetOrdersUseCase struct {
	port gateways.GetOrdersPort
}

func NewGetOrdersUseCase(port gateways.GetOrdersPort) *GetOrdersUseCase {
	return &GetOrdersUseCase{
		port: port,
	}
}

func (itself *GetOrdersUseCase) Execute(pagination global.Pagination, listeners contracts.GetOrdersListeners) {
	orders, err := itself.port.Execute(pagination)
	if err != nil {
		listeners.OnError(err)
		return
	}
	listeners.OnSuccess(orders)
}
