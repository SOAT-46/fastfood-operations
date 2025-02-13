//go:build unit

package gateways_test

import (
	"context"
	"errors"
	"testing"

	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/gateways"
	"github.com/SOAT-46/fastfood-operations/test/orders/repositories/builders"
	"github.com/SOAT-46/fastfood-operations/test/orders/repositories/doubles"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ErrTestGetOrdersGateway = errors.New("test error")

func TestGetOrdersGateway(t *testing.T) {
	paginatedEntity := builders.NewGormOrderBuilder().BuildPaginated()
	t.Run("should return a list of orders", func(t *testing.T) {
		// given
		repository := doubles.NewInMemoryOrdersRepository().WithOrders(paginatedEntity)
		gateway := gateways.NewGetOrdersGateway(repository)

		// when
		orders, err := gateway.Execute(context.Background(), paginatedEntity.Pagination)

		// then
		assert.NotNil(t, orders)
		require.NoError(t, err)
	})

	t.Run("should return an error to get the orders", func(t *testing.T) {
		// given
		repository := doubles.NewInMemoryOrdersRepository().WithError(ErrTestGetOrdersGateway)
		gateway := gateways.NewGetOrdersGateway(repository)

		// when
		_, err := gateway.Execute(context.Background(), paginatedEntity.Pagination)

		// then
		require.Error(t, err)
	})
}
