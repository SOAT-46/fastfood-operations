//go:build unit

package implementations_test

import (
	"testing"

	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/contracts"
	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/implementations"
	entities2 "github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	"github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
	builders2 "github.com/SOAT-46/fastfood-operations/test/orders/domain/builders"
	"github.com/SOAT-46/fastfood-operations/test/orders/gateways/doubles"
	"github.com/SOAT-46/fastfood-operations/test/shared/builders"
	"github.com/stretchr/testify/assert"
)

func TestGetOrdersUseCase(t *testing.T) {
	t.Run("should call OnSuccess when there are orders", func(t *testing.T) {
		// given
		pagination := builders.NewPaginationBuilder().Build()
		order := builders2.NewOrderBuilder().BuildPaginated()

		port := doubles.NewInMemoryGetOrdersPort().WithOrders(order)
		useCase := implementations.NewGetOrdersUseCase(port)

		listeners := contracts.GetOrdersListeners{
			OnSuccess: func(orders entities.PaginatedEntity[entities2.Order]) {
				// then
				assert.NotNil(t, orders)
			},
		}

		// when
		useCase.Execute(pagination, listeners)
	})

	t.Run("should call OnError when there's an error", func(t *testing.T) {
		// given
		port := doubles.NewInMemoryGetOrdersPort().WithError()
		useCase := implementations.NewGetOrdersUseCase(port)

		listeners := contracts.GetOrdersListeners{
			OnError: func(err error) {
				// then
				assert.Error(t, err)
			},
		}

		// when
		useCase.Execute(builders.NewPaginationBuilder().Build(), listeners)
	})
}
