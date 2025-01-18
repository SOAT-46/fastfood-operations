//go:build unit

package controllers_test

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/http/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateOrderController(t *testing.T) {
	endpoint := "/v1/orders"
	t.Run("should return status code CREATED (201)", func(t *testing.T) {
		// given
		gin.SetMode(gin.TestMode)

		controller := controllers.NewCreateOrderController()
		router := gin.New()

		// when
		router.POST(endpoint, controller.Execute)

		req, err := http.NewRequest(http.MethodPost, endpoint, nil)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// then
		assert.NoError(t, err, "no error in the request")
		assert.Equal(t, http.StatusCreated, recorder.Code, "status code is 201")
	})
}
