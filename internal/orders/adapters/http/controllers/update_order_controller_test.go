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
	t.Run("should respond OK (200) when update the order", func(t *testing.T) {
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
		assert.NoError(t, err, defaultNoErrorMessage)
	})

	t.Run("should respond BAD_REQUEST (400) when the order was invalid", func(t *testing.T) {
		// given
		gin.SetMode(gin.TestMode)
		router := gin.New()

		entity := builders.NewUpdateOrderRequestBuilder().BuildInvalidRequest()
		useCase := doubles.NewUpdateOrderUseCaseStub().WithOnSuccess()
		controller := controllers.NewUpdateOrderController(useCase)

		// when
		router.PUT(endpoint, controller.Execute)
		req, err := http.NewRequest(http.MethodPut, endpoint, entity)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// then
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		assert.NoError(t, err, defaultNoErrorMessage)
	})

	t.Run("should respond NOT_FOUND (404) when the order was not found", func(t *testing.T) {
		// given
		gin.SetMode(gin.TestMode)
		router := gin.New()

		entity := builders.NewUpdateOrderRequestBuilder().BuildRequest()
		useCase := doubles.NewUpdateOrderUseCaseStub().WithOnNotFound()
		controller := controllers.NewUpdateOrderController(useCase)

		// when
		router.PUT(endpoint, controller.Execute)
		req, err := http.NewRequest(http.MethodPut, endpoint, entity)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// then
		assert.Equal(t, http.StatusNotFound, recorder.Code)
		assert.NoError(t, err, defaultNoErrorMessage)
	})

	t.Run("should respond INTERNAL_SERVER_ERROR (500) due to an unexpected error", func(t *testing.T) {
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
		assert.NoError(t, err, defaultNoErrorMessage)
	})

	t.Run("should return the metadata", func(t *testing.T) {
		// given
		controller := controllers.NewUpdateOrderController(nil)

		// when
		metadata := controller.GetBind()

		// then
		assert.Equal(t, "/orders/:id", metadata.RelativePath)
		assert.Equal(t, http.MethodPut, metadata.Method)
	})
}
