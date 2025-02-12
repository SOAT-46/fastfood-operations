//go:build unit

package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/http/controllers"
	"github.com/SOAT-46/fastfood-operations/test/orders/application/usecases/doubles"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetOrdersController(t *testing.T) {
	endpoint := "/orders"
	t.Run("should respond OK (200) when there are orders to show", func(t *testing.T) {
		// given
		gin.SetMode(gin.TestMode)
		router := gin.New()

		useCase := doubles.NewGetOrdersUseCaseStub().WithOnSuccess()
		controller := controllers.NewGetOrdersController(useCase)

		// when
		router.GET(endpoint, controller.Execute)

		req, err := http.NewRequest(http.MethodGet, endpoint, nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// then
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.NoError(t, err, defaultNoErrorMessage)
	})

	t.Run("should respond INTERNAL_SERVER_ERROR (500) due to an unexpected error", func(t *testing.T) {
		// given
		gin.SetMode(gin.TestMode)
		router := gin.New()

		useCase := doubles.NewGetOrdersUseCaseStub().WithOnError()
		controller := controllers.NewGetOrdersController(useCase)

		// when
		router.GET(endpoint, controller.Execute)

		req, err := http.NewRequest(http.MethodGet, endpoint, nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// then
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
		assert.NoError(t, err, defaultNoErrorMessage)
	})

	t.Run("should return the metadata", func(t *testing.T) {
		// given
		controller := controllers.NewGetOrdersController(nil)

		// when
		metadata := controller.GetBind()

		// then
		assert.Equal(t, "/orders", metadata.RelativePath)
		assert.Equal(t, http.MethodGet, metadata.Method)
	})
}
