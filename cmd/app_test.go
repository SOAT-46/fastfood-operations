//go:build unit

package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	cmd "github.com/SOAT-46/fastfood-operations/cmd"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCorsMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(cmd.CorsMiddleware())
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	t.Run("should set CORS headers", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, req)

		assert.Equal(t, "*", recorder.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, "GET, POST, PUT, DELETE, PATCH, OPTIONS", recorder.Header().Get("Access-Control-Allow-Methods"))
		assert.Equal(t, "Content-Type, Authorization", recorder.Header().Get("Access-Control-Allow-Headers"))
	})

	t.Run("should handle OPTIONS request", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodOptions, "/test", nil)
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
