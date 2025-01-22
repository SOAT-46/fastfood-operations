//go:build unit

package implementations_test

import (
	"testing"

	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/contracts"
	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/implementations"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/exceptions"
	"github.com/SOAT-46/fastfood-operations/test/orders/domain/builders"
	"github.com/SOAT-46/fastfood-operations/test/orders/gateways/doubles"
	"github.com/stretchr/testify/assert"
)

func TestUpdateOrderUseCase(t *testing.T) {
	t.Run("should call OnSuccess when the order was updated", func(t *testing.T) {
		// given
		order := builders.NewOrderBuilder().Build()

		port := doubles.NewInMemoryUpdateOrderPort().WithOrder(&order)
		getPort := doubles.NewInMemoryGetOrderByIDPort().WithOrder(&order)
		useCase := implementations.NewUpdateOrderUseCase(port, getPort)

		listeners := contracts.UpdateOrderListeners{
			OnSuccess: func(order entities.Order) {
				// then
				assert.NotNil(t, order)
			},
		}

		// when
		useCase.Execute(order, listeners)
	})

	t.Run("should call OnNotFound when there's an unexpected error", func(t *testing.T) {
		// given
		order := builders.NewOrderBuilder().Build()

		port := doubles.NewInMemoryUpdateOrderPort()
		getPort := doubles.NewInMemoryGetOrderByIDPort().WithError(exceptions.ErrOrderNotFound)
		useCase := implementations.NewUpdateOrderUseCase(port, getPort)

		listeners := contracts.UpdateOrderListeners{
			OnError: func(err error) {
				// then
				assert.Error(t, err)
			},
		}

		// when
		useCase.Execute(order, listeners)
	})
}
