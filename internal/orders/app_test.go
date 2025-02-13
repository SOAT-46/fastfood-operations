//go:build unit

package orders_test

import (
	"testing"

	"github.com/SOAT-46/fastfood-operations/internal/orders"
	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	t.Run("should create the app", func(t *testing.T) {
		// given
		app := orders.NewApp(nil, nil, nil, nil)

		// when
		controllers := app.GetControllers()

		// then
		assert.Len(t, controllers, 4)
	})
}
