package doubles

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/repositories/models"
	"github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
)

type InMemoryOrdersRepository struct {
	err    error
	orders entities.PaginatedEntity[models.GormOrder]
}

func NewInMemoryOrdersRepository() *InMemoryOrdersRepository {
	return &InMemoryOrdersRepository{}
}

func (itself *InMemoryOrdersRepository) WithOrders(
	orders entities.PaginatedEntity[models.GormOrder]) *InMemoryOrdersRepository {
	itself.orders = orders
	return itself
}

func (itself *InMemoryOrdersRepository) WithError(err error) *InMemoryOrdersRepository {
	itself.err = err
	return itself
}

func (itself *InMemoryOrdersRepository) ListAll(
	_ entities.Pagination) (entities.PaginatedEntity[models.GormOrder], error) {
	return itself.orders, itself.err
}

func (itself *InMemoryOrdersRepository) GetByID(_ int) (*models.GormOrder, error) {
	if itself.err != nil {
		return nil, itself.err
	}
	return &itself.orders.Content[0], itself.err
}

func (itself *InMemoryOrdersRepository) Save(order models.GormOrder) (*models.GormOrder, error) {
	return &order, itself.err
}

func (itself *InMemoryOrdersRepository) Update(order models.GormOrder) (*models.GormOrder, error) {
	return &order, itself.err
}
