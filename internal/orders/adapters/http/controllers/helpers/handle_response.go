package helpers

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/http/controllers/mappers"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	global "github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
	"github.com/gin-gonic/gin"
)

func HandlePaginatedResponse(
	ctx *gin.Context,
	statusCode int,
	page global.PaginatedEntity[entities.Order]) {
	response := mappers.MapToPaginatedResponse(page)
	ctx.JSON(statusCode, response)
}

func HandleResponse(
	ctx *gin.Context,
	statusCode int,
	order entities.Order) {
	response := mappers.MapToResponse(order)
	ctx.JSON(statusCode, response)
}
