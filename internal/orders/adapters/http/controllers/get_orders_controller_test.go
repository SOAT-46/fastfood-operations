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
	t.Run("should get a list of orders", func(t *testing.T) {
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
		assert.NoError(t, err, "no error in the request")
	})

	t.Run("should return an error to get orders", func(t *testing.T) {
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
		assert.NoError(t, err, "no error in the request")
	})
}
