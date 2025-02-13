package doubles

import (
	"context"

	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/repositories/models"
	"github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
)

type InMemoryOrdersRepository struct {
	err    error
	orders entities.PaginatedEntity[models.MongoOrder]
}

func NewInMemoryOrdersRepository() *InMemoryOrdersRepository {
	return &InMemoryOrdersRepository{}
}

func (itself *InMemoryOrdersRepository) WithOrders(
	orders entities.PaginatedEntity[models.MongoOrder]) *InMemoryOrdersRepository {
	itself.orders = orders
	return itself
}

func (itself *InMemoryOrdersRepository) WithError(err error) *InMemoryOrdersRepository {
	itself.err = err
	return itself
}

func (itself *InMemoryOrdersRepository) ListAll(
	_ context.Context, _ entities.Pagination) (entities.PaginatedEntity[models.MongoOrder], error) {
	return itself.orders, itself.err
}

func (itself *InMemoryOrdersRepository) GetByID(_ context.Context, _ string) (*models.MongoOrder, error) {
	if itself.err != nil {
		return nil, itself.err
	}
	return &itself.orders.Content[0], itself.err
}

func (itself *InMemoryOrdersRepository) Save(_ context.Context, order models.MongoOrder) (*models.MongoOrder, error) {
	return &order, itself.err
}

func (itself *InMemoryOrdersRepository) Update(_ context.Context, order models.MongoOrder) (*models.MongoOrder, error) {
	return &order, itself.err
}
