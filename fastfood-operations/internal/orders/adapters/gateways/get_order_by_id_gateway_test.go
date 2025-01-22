package gateways_test

import (
  "errors"
  "github.com/SOAT-46/fastfood-operations/internal/orders/adapters/gateways"
  "github.com/SOAT-46/fastfood-operations/test/orders/repositories/builders"
  "github.com/SOAT-46/fastfood-operations/test/orders/repositories/doubles"
  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/require"
  "testing"
)

var ErrTestGetOrderByIDGateway = errors.New("test error")

func TestGetOrderByIDGateway(t *testing.T) {
  id := 1
  t.Run("should return an order by id", func(t *testing.T) {
    // given
    paginatedEntity := builders.NewGormOrderBuilder().BuildPaginated()
    repository := doubles.NewInMemoryOrdersRepository().WithOrders(paginatedEntity)
    gateway := gateways.NewGetOrderByIDGateway(repository)

    // when
    order, err := gateway.Execute(id)

    // then
    assert.NotNil(t, order)
    require.NoError(t, err)
  })

  t.Run("should return an error to get the order by id", func(t *testing.T) {
    // given
    repository := doubles.NewInMemoryOrdersRepository().WithError(ErrTestGetOrderByIDGateway)
    gateway := gateways.NewGetOrderByIDGateway(repository)

    // when
    order, err := gateway.Execute(id)

    // then
    assert.Nil(t, order)
    require.Error(t, err)
  })
}
