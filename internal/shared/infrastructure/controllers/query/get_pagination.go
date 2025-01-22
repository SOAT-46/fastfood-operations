package query

import (
	"strconv"

	"github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
	"github.com/gin-gonic/gin"
)

func GetPagination(gcontext *gin.Context) entities.Pagination {
	pagination := entities.Pagination{}

	pagination.Page, _ = strconv.Atoi(gcontext.DefaultQuery("page", "1"))
	pagination.Size, _ = strconv.Atoi(gcontext.DefaultQuery("size", "10"))
	pagination.Filter = gcontext.Query("filter")
	return pagination
}
