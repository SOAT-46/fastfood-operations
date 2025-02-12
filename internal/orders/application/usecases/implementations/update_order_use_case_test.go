//go:build unit

package implementations_test

import (
	"context"
	"testing"

	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/contracts"
	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/implementations"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/exceptions"
	doubles2 "github.com/SOAT-46/fastfood-operations/test/orders/application/usecases/doubles"
	"github.com/SOAT-46/fastfood-operations/test/orders/domain/builders"
	"github.com/SOAT-46/fastfood-operations/test/orders/gateways/doubles"
	"github.com/stretchr/testify/assert"
)

func TestUpdateOrderUseCase(t *testing.T) {
	t.Run("should call OnSuccess when the order was updated", func(t *testing.T) {
		// given
		order := builders.NewOrderBuilder().Build()

		port := doubles.NewInMemoryUpdateOrderPort().WithOrder(&order)
		getPort := doubles2.NewGetOrderByIDUseCaseStub().WithOnSuccess()
		useCase := implementations.NewUpdateOrderUseCase(port, getPort)

		listeners := contracts.UpdateOrderListeners{
			OnSuccess: func(order entities.Order) {
				// then
				assert.NotNil(t, order)
			},
		}

		// when
		useCase.Execute(context.Background(), order, listeners)
	})

	t.Run("should call OnError when there's an error to update the order", func(t *testing.T) {
		// given
		order := builders.NewOrderBuilder().Build()

		port := doubles.NewInMemoryUpdateOrderPort().WithError()
		getPort := doubles2.NewGetOrderByIDUseCaseStub().WithOnSuccess()
		useCase := implementations.NewUpdateOrderUseCase(port, getPort)

		listeners := contracts.UpdateOrderListeners{
			OnError: func(err error) {
				// then
				assert.Error(t, err)
			},
		}

		// when
		useCase.Execute(context.Background(), order, listeners)
	})

	t.Run("should call OnNotFound when there's an unexpected error", func(t *testing.T) {
		// given
		order := builders.NewOrderBuilder().Build()

		port := doubles.NewInMemoryUpdateOrderPort()
		getPort := doubles2.NewGetOrderByIDUseCaseStub().WithOnError(exceptions.ErrOrderNotFound)
		useCase := implementations.NewUpdateOrderUseCase(port, getPort)

		listeners := contracts.UpdateOrderListeners{
			OnError: func(err error) {
				// then
				assert.Error(t, err)
			},
		}

		// when
		useCase.Execute(context.Background(), order, listeners)
	})
}
