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

func TestCreateOrderController(t *testing.T) {
	endpoint := "/orders"
	t.Run("should create an order", func(t *testing.T) {
		// given
		gin.SetMode(gin.TestMode)
		router := gin.New()

		entity := builders.NewCreateOrderRequestBuilder().BuildRequest()
		useCase := doubles.NewCreateOrderUseCaseStub().WithOnSuccess()
		controller := controllers.NewCreateOrderController(useCase)

		// when
		router.POST(endpoint, controller.Execute)

		req, err := http.NewRequest(http.MethodPost, endpoint, entity)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// then
		assert.Equal(t, http.StatusCreated, recorder.Code)
		assert.NoError(t, err, "no error in the request")
	})

	t.Run("should return an error to create the order", func(t *testing.T) {
		// given
		gin.SetMode(gin.TestMode)
		router := gin.New()

		entity := builders.NewCreateOrderRequestBuilder().BuildRequest()
		useCase := doubles.NewCreateOrderUseCaseStub().WithOnError()
		controller := controllers.NewCreateOrderController(useCase)

		// when
		router.POST(endpoint, controller.Execute)

		req, err := http.NewRequest(http.MethodPost, endpoint, entity)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// then
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
		assert.NoError(t, err, "no error in the request")
	})
}
