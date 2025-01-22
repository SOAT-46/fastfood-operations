package gateways_test

import (
	"errors"
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/gateways"
	"github.com/SOAT-46/fastfood-operations/test/orders/domain/builders"
	"github.com/SOAT-46/fastfood-operations/test/orders/repositories/doubles"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ErrTestUpdateOrderGateway = errors.New("test error")

func TestUpdateOrderGateway(t *testing.T) {
	t.Run("should update an order", func(t *testing.T) {
		// given
		entity := builders.NewOrderBuilder().Build()
		repository := doubles.NewInMemoryOrdersRepository()
		gateway := gateways.NewUpdateOrderGateway(repository)

		// when
		order, err := gateway.Execute(entity)

		// then
		assert.NotNil(t, order)
		assert.NoError(t, err)
	})

	t.Run("should return an error to update the order", func(t *testing.T) {
		// given
		entity := builders.NewOrderBuilder().Build()
		repository := doubles.NewInMemoryOrdersRepository().WithError(ErrTestUpdateOrderGateway)
		gateway := gateways.NewUpdateOrderGateway(repository)

		// when
		order, err := gateway.Execute(entity)

		// then
		assert.Nil(t, order)
		assert.Error(t, err)
	})
}
