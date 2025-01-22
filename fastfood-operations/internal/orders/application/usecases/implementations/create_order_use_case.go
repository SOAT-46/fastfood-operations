package implementations

import (
	"fmt"
	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/contracts"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/gateways"
)

type CreateOrderUseCase struct {
	port gateways.SaveOrderPort
}

func NewCreateOrderUseCase(port gateways.SaveOrderPort) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		port: port,
	}
}

func (itself *CreateOrderUseCase) Execute(
	input entities.CreateOrderInput, listeners contracts.CreateOrderListeners) {
	if !input.IsValid() {
		listeners.OnInvalid(fmt.Errorf("invalid order, impossible to save"))
		return
	}

	order, err := itself.port.Execute(input)
	if err != nil {
		listeners.OnError(err)
		return
	}
	listeners.OnSuccess(*order)
}
