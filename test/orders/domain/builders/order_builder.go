package builders

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	entities2 "github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
	"github.com/SOAT-46/fastfood-operations/test/shared/builders"
)

type OrderBuilder struct {
	builders.BaseBuilder[entities.Order]
}

func NewOrderBuilder() *OrderBuilder {
	return &OrderBuilder{}
}

func (itself *OrderBuilder) BuildPaginated() entities2.PaginatedEntity[entities.Order] {
	data := itself.BuildMany()
	pagination := builders.NewPaginationBuilder().Build()

	return entities2.NewPaginatedEntity(data, pagination)
}
