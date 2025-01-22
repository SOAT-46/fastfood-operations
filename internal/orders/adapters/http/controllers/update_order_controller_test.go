//go:build unit

package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/http/controllers"
	"github.com/SOAT-46/fastfood-operations/test/orders/adapters/http/requests/builders"
	"github.com/SOAT-46/fastfood-operations/test/orders/application/usecases/doubles"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUpdateOrderController(t *testing.T) {
	endpoint := "/orders/:id"
	t.Run("should update an order", func(t *testing.T) {
		// given
		gin.SetMode(gin.TestMode)
		router := gin.New()

		entity := builders.NewUpdateOrderRequestBuilder().BuildRequest()
		useCase := doubles.NewUpdateOrderUseCaseStub().WithOnSuccess()
		controller := controllers.NewUpdateOrderController(useCase)

		// when
		router.PUT(endpoint, controller.Execute)
		req, err := http.NewRequest(http.MethodPut, endpoint, entity)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// then
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.NoError(t, err, "no error in the request")
	})

	t.Run("should return an error to update the order", func(t *testing.T) {
		// given
		gin.SetMode(gin.TestMode)
		router := gin.New()

		entity := builders.NewUpdateOrderRequestBuilder().BuildRequest()
		useCase := doubles.NewUpdateOrderUseCaseStub().WithOnError()
		controller := controllers.NewUpdateOrderController(useCase)

		// when
		router.PUT(endpoint, controller.Execute)
		req, err := http.NewRequest(http.MethodPut, endpoint, entity)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// then
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
		assert.NoError(t, err, "no error in the request")
	})
}
