package contracts

import (
	"context"

	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/repositories/models"
	"github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
)

type OrdersRepository interface {
	ListAll(ctx context.Context, pagination entities.Pagination) (entities.PaginatedEntity[models.MongoOrder], error)
	GetByID(ctx context.Context, id string) (*models.MongoOrder, error)
	Save(ctx context.Context, order models.MongoOrder) (*models.MongoOrder, error)
	Update(ctx context.Context, order models.MongoOrder) (*models.MongoOrder, error)
}
