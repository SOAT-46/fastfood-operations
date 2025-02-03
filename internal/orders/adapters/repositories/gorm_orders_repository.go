package repositories

import (
	"errors"
	"fmt"

	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/repositories/models"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/exceptions"
	"github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
	"gorm.io/gorm"
)

type GormOrdersRepository struct {
	client *gorm.DB
}

func NewGormOrdersRepository(client *gorm.DB) *GormOrdersRepository {
	return &GormOrdersRepository{
		client: client,
	}
}

func (itself *GormOrdersRepository) ListAll(
	pagination entities.Pagination) (entities.PaginatedEntity[models.GormOrder], error) {
	where := "status IN ('RECEIVED', 'PREPARATION', 'READY')"
	var list []models.GormOrder

	query := itself.client.Model(&list).Where(where)

	var totalElements int64
	err := query.
		Model(&models.GormOrder{}).
		Where(where).
		Count(&totalElements).
		Error
	if err != nil {
		return entities.PaginatedEntity[models.GormOrder]{},
			fmt.Errorf("can't count in orders repository. Reason: %w", err)
	}
	pagination.TotalElements = totalElements

	err = query.
		Where(where).
		Order("status DESC").
		Order("received_at ASC").
		Offset(pagination.Offset()).
		Limit(pagination.Size).
		Find(&list).
		Error
	if err != nil {
		return entities.PaginatedEntity[models.GormOrder]{},
			fmt.Errorf("can't ListAll in orders repository. Reason: %w", err)
	}

	return entities.NewPaginatedEntity(list, pagination), nil
}

func (itself *GormOrdersRepository) GetByID(id int) (*models.GormOrder, error) {
	order := &models.GormOrder{}

	err := itself.
		client.
		Preload("Items").
		Where("id = ?", id).
		First(order).
		Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, exceptions.ErrOrderNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("can't GetByID in orders repository. Reason: %w", err)
	}
	return order, nil
}

func (itself *GormOrdersRepository) Save(order models.GormOrder) (*models.GormOrder, error) {
	err := itself.client.Save(&order).Error
	if err != nil {
		return nil, fmt.Errorf("can't Save in orders repository. Reason: %w", err)
	}
	return &order, nil
}

func (itself *GormOrdersRepository) Update(order models.GormOrder) (*models.GormOrder, error) {
	err := itself.client.Updates(&order).Error
	if err != nil {
		return nil, fmt.Errorf("can't Update in orders repository. Reason: %w", err)
	}
	return &order, nil
}
