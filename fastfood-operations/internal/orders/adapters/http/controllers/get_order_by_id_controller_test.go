package controllers_test

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/http/controllers"
	"github.com/SOAT-46/fastfood-operations/test/orders/application/usecases/doubles"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetOrderByIDController(t *testing.T) {
	endpoint := "/orders/:id"
	t.Run("should get an order by id", func(t *testing.T) {
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
		assert.NoError(t, err, "no error in the request")
	})

	t.Run("should return an error to get the order by id", func(t *testing.T) {
		// given
		gin.SetMode(gin.TestMode)
		router := gin.New()

		useCase := doubles.NewGetOrderByIDUseCaseStub().WithOnError()
		controller := controllers.NewGetOrderByIDController(useCase)

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
