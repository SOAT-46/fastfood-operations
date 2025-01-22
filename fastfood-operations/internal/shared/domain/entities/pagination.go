package entities

type Pagination struct {
	Page          int
	Size          int
	TotalElements int64
	Filter        string
}

func NewPagination(page int, size int, totalElements int64) Pagination {
	return Pagination{
		Page:          page,
		Size:          size,
		TotalElements: totalElements,
	}
}

func (itself *Pagination) FinalIndex() int {
	return itself.Page*itself.Size - 1
}

func (itself *Pagination) Offset() int {
	return (itself.Page - 1) * itself.Size
}

func (itself *Pagination) IsFirstPage() bool {
	return itself.Page == 1
}

func (itself *Pagination) IsLastPage() bool {
	return int64(itself.Page*itself.Size) >= itself.TotalElements
}

func (itself *Pagination) TotalPages() int64 {
	limit := int64(itself.Size)

	if limit >= itself.TotalElements {
		return 1
	}
	return itself.TotalElements / limit
}
