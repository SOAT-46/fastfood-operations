package builders

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/repositories/models"
	"github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
	"github.com/SOAT-46/fastfood-operations/test/shared/builders"
)

type GormOrderBuilder struct {
	builders.BaseBuilder[models.GormOrder]
}

func NewGormOrderBuilder() *GormOrderBuilder {
	return &GormOrderBuilder{}
}

func (itself *GormOrderBuilder) BuildPaginated() entities.PaginatedEntity[models.GormOrder] {
	data := itself.BuildMany()
	pagination := builders.NewPaginationBuilder().Build()

	return entities.NewPaginatedEntity(data, pagination)
}
