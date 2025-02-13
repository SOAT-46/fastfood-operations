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

const (
	defaultNoErrorMessage = "no error in the request"
)

func TestCreateOrderController(t *testing.T) {
	endpoint := "/orders"
	t.Run("should respond CREATED (201) when the order was created", func(t *testing.T) {
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
		assert.NoError(t, err, defaultNoErrorMessage)
	})

	t.Run("should respond BAD_REQUEST (400) when the request is invalid", func(t *testing.T) {
		// given
		gin.SetMode(gin.TestMode)
		router := gin.New()

		entity := builders.NewCreateOrderRequestBuilder().BuildInvalidRequest()
		useCase := doubles.NewCreateOrderUseCaseStub().WithOnSuccess()
		controller := controllers.NewCreateOrderController(useCase)

		// when
		router.POST(endpoint, controller.Execute)

		req, err := http.NewRequest(http.MethodPost, endpoint, entity)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// then
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		assert.NoError(t, err, defaultNoErrorMessage)
	})

	t.Run("should respond INTERNAL_SERVER_ERROR (500) when there's an error to create the order",
		func(t *testing.T) {
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
			assert.NoError(t, err, defaultNoErrorMessage)
		})

	t.Run("should return the metadata", func(t *testing.T) {
		// given
		controller := controllers.NewCreateOrderController(nil)

		// when
		metadata := controller.GetBind()

		// then
		assert.Equal(t, "/orders", metadata.RelativePath)
		assert.Equal(t, http.MethodPost, metadata.Method)
	})
}
