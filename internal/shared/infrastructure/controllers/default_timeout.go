package controllers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	timeoutSeconds = 5
)

func DefaultTimeout(gcontext *gin.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(gcontext.Request.Context(), timeoutSeconds*time.Second)
	return ctx, cancel
}
