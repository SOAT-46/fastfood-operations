package mappers

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/repositories/models"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	global "github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
)

func MapToPaginatedEntity(dbPage global.PaginatedEntity[models.GormOrder]) global.PaginatedEntity[entities.Order] {
	var orders []entities.Order
	for _, order := range dbPage.Content {
		mapped, err := order.ToDomain()
		if err != nil {
			break
		}
		orders = append(orders, *mapped)
	}
	return global.NewPaginatedEntity(orders, dbPage.Pagination)
}
