//go:build unit

package implementations_test

import (
	"testing"

	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/contracts"
	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/implementations"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	"github.com/SOAT-46/fastfood-operations/test/orders/domain/builders"
	"github.com/SOAT-46/fastfood-operations/test/orders/gateways/doubles"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrderUseCase(t *testing.T) {
	t.Run("should call OnSuccess when the order was saved", func(t *testing.T) {
		// given
		input := builders.NewCreateOrderInputBuilder().Build()
		order := builders.NewOrderBuilder().Build()

		port := doubles.NewInMemorySaveOrderPort().WithOrder(&order)
		useCase := implementations.NewCreateOrderUseCase(port)

		listeners := contracts.CreateOrderListeners{
			OnSuccess: func(order entities.Order) {
				// then
				assert.NotNil(t, order)
			},
		}

		// when
		useCase.Execute(input, listeners)
	})

	t.Run("should call OnInvalid when the input is not valid", func(t *testing.T) {
		// given
		input := builders.NewCreateOrderInputBuilder().BuildInvalid()

		port := doubles.NewInMemorySaveOrderPort()
		useCase := implementations.NewCreateOrderUseCase(port)

		listeners := contracts.CreateOrderListeners{
			OnInvalid: func(err error) {
				// then
				assert.Error(t, err)
			},
		}

		// when
		useCase.Execute(input, listeners)
	})

	t.Run("should call OnError when there`s an error to save the order", func(t *testing.T) {
		// given
		input := builders.NewCreateOrderInputBuilder().Build()

		port := doubles.NewInMemorySaveOrderPort().WithError()
		useCase := implementations.NewCreateOrderUseCase(port)

		listeners := contracts.CreateOrderListeners{
			OnError: func(err error) {
				// then
				assert.Error(t, err)
			},
		}

		// when
		useCase.Execute(input, listeners)
	})
}
