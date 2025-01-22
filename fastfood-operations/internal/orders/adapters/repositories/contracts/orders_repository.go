package contracts

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/repositories/models"
	"github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
)

type OrdersRepository interface {
	ListAll(pagination entities.Pagination) (entities.PaginatedEntity[models.GormOrder], error)
	GetByID(id int) (*models.GormOrder, error)
	Save(order models.GormOrder) (*models.GormOrder, error)
	Update(order models.GormOrder) (*models.GormOrder, error)
}
