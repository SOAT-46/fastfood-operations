//go:build unit

package implementations_test

import (
	"errors"
	"testing"

	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/contracts"
	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/implementations"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/exceptions"
	"github.com/SOAT-46/fastfood-operations/test/orders/domain/builders"
	"github.com/SOAT-46/fastfood-operations/test/orders/gateways/doubles"
	"github.com/stretchr/testify/assert"
)

var ErrTestGetOrderByIDUseCase = errors.New("test error")

func TestGetOrderByIDUseCase(t *testing.T) {
	id := 1
	t.Run("should call OnSuccess when the order was found", func(t *testing.T) {
		// given
		order := builders.NewOrderBuilder().Build()
		port := doubles.NewInMemoryGetOrderByIDPort().WithOrder(&order)
		useCase := implementations.NewGetOrderByIDUseCase(port)

		listeners := contracts.GetOrderByIDListeners{
			OnSuccess: func(order entities.Order) {
				// then
				assert.NotNil(t, order)
			},
		}

		// when
		useCase.Execute(id, listeners)
	})

	t.Run("should call OnNotFound when the order was not found", func(t *testing.T) {
		// given
		var wasCalled = false
		port := doubles.NewInMemoryGetOrderByIDPort().WithError(exceptions.ErrOrderNotFound)
		useCase := implementations.NewGetOrderByIDUseCase(port)

		listeners := contracts.GetOrderByIDListeners{
			OnNotFound: func() {
				wasCalled = true
			},
		}

		// when
		useCase.Execute(id, listeners)

		// then
		assert.True(t, wasCalled)
	})

	t.Run("should call OnError when there's an exception to get the order", func(t *testing.T) {
		// given
		port := doubles.NewInMemoryGetOrderByIDPort().
			WithError(ErrTestGetOrderByIDUseCase)
		useCase := implementations.NewGetOrderByIDUseCase(port)

		listeners := contracts.GetOrderByIDListeners{
			OnError: func(err error) {
				// then
				assert.Error(t, err)
			},
		}

		// when
		useCase.Execute(id, listeners)
	})
}
