//go:build unit

package gateways_test

import (
	"errors"
	"testing"

	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/gateways"
	"github.com/SOAT-46/fastfood-operations/test/orders/domain/builders"
	"github.com/SOAT-46/fastfood-operations/test/orders/repositories/doubles"
	"github.com/stretchr/testify/assert"
)

var ErrTestSaveOrderGateway = errors.New("test error")

func TestSaveOrderGateway(t *testing.T) {
	t.Run("should save an order", func(t *testing.T) {
		// given
		repository := doubles.NewInMemoryOrdersRepository()
		gateway := gateways.NewSaveOrderGateway(repository)

		// when
		order, err := gateway.Execute(builders.NewCreateOrderInputBuilder().Build())

		// then
		assert.NotNil(t, order)
		assert.NoError(t, err)
	})

	t.Run("should return an error to save the order", func(t *testing.T) {
		// given
		repository := doubles.NewInMemoryOrdersRepository().WithError(ErrTestSaveOrderGateway)
		gateway := gateways.NewSaveOrderGateway(repository)

		// when
		order, err := gateway.Execute(builders.NewCreateOrderInputBuilder().Build())

		// then
		assert.Nil(t, order)
		assert.Error(t, err)
	})
}
