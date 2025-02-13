package builders

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/repositories/models"
	"github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
	"github.com/SOAT-46/fastfood-operations/test/shared/builders"
)

type GormOrderBuilder struct {
	builders.BaseBuilder[models.MongoOrder]
}

func NewGormOrderBuilder() *GormOrderBuilder {
	return &GormOrderBuilder{}
}

func (itself *GormOrderBuilder) BuildPaginated() entities.PaginatedEntity[models.MongoOrder] {
	data := itself.BuildMany()
	pagination := builders.NewPaginationBuilder().Build()

	return entities.NewPaginatedEntity(data, pagination)
}
