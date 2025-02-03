package builders

import global "github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"

type PaginationBuilder struct {
	BaseBuilder[global.Pagination]
}

func NewPaginationBuilder() *PaginationBuilder {
	return &PaginationBuilder{}
}
