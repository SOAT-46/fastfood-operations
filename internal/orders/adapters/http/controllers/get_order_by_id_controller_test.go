//go:build unit

package controllers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/http/controllers"
	"github.com/SOAT-46/fastfood-operations/test/orders/application/usecases/doubles"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ErrTestGetOrderByIDController = errors.New("test error")

func TestGetOrderByIDController(t *testing.T) {
	endpoint := "/orders/:id"
	t.Run("should respond OK (200) when the order was found", func(t *testing.T) {
		// given
		gin.SetMode(gin.TestMode)
		router := gin.New()

		useCase := doubles.NewGetOrderByIDUseCaseStub().WithOnSuccess()
		controller := controllers.NewGetOrderByIDController(useCase)

		// when
		router.GET(endpoint, controller.Execute)

		req, err := http.NewRequest(http.MethodGet, endpoint, nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// then
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.NoError(t, err, defaultNoErrorMessage)
	})

	t.Run("should respond NOT_FOUND (404) when the order was not found", func(t *testing.T) {
		// given
		gin.SetMode(gin.TestMode)
		router := gin.New()

		useCase := doubles.NewGetOrderByIDUseCaseStub().WithOnNotFound()
		controller := controllers.NewGetOrderByIDController(useCase)

		// when
		router.GET(endpoint, controller.Execute)

		req, err := http.NewRequest(http.MethodGet, endpoint, nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// then
		assert.Equal(t, http.StatusNotFound, recorder.Code)
		assert.NoError(t, err, defaultNoErrorMessage)
	})

	t.Run("should respond INTERNAL_SERVER_ERROR (500) when the order was not found", func(t *testing.T) {
		// given
		gin.SetMode(gin.TestMode)
		router := gin.New()

		useCase := doubles.NewGetOrderByIDUseCaseStub().
			WithOnError(ErrTestGetOrderByIDController)
		controller := controllers.NewGetOrderByIDController(useCase)

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
		controller := controllers.NewGetOrderByIDController(nil)

		// when
		metadata := controller.GetBind()

		// then
		assert.Equal(t, "/orders/:id", metadata.RelativePath)
		assert.Equal(t, http.MethodGet, metadata.Method)
	})
}
