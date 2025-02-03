package helpers

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/http/mappers"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	global "github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
	"github.com/gin-gonic/gin"
)

func HandlePaginatedResponse(
	ctx *gin.Context,
	statusCode int,
	page global.PaginatedEntity[entities.Order],
	onError func(err error)) {
	response, err := mappers.MapToPaginatedResponse(page)
	if err != nil {
		onError(err)
	}
	ctx.JSON(statusCode, response)
}

func HandleResponse(
	ctx *gin.Context,
	statusCode int,
	order entities.Order,
	onError func(err error)) {
	response, err := mappers.MapToResponse(order)
	if err != nil {
		onError(err)
	}
	ctx.JSON(statusCode, response)
}
