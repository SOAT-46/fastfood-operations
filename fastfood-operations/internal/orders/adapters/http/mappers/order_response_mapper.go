package mappers

import (
	"fmt"
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/http/responses"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	global "github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
	responses2 "github.com/SOAT-46/fastfood-operations/internal/shared/infrastructure/controllers/page"
	"github.com/mitchellh/mapstructure"
)

func MapToResponse(order entities.Order) (*responses.OrderResponse, error) {
	var response responses.OrderResponse
	err := mapstructure.Decode(order, &response)
	if err != nil {
		return nil, fmt.Errorf("can't decode MapToResponse. Reason: %w", err)
	}
	return &response, nil
}

func MapToResponses(orders []entities.Order) ([]responses.OrderResponse, error) {
	var orderResponses []responses.OrderResponse

	for _, order := range orders {
		mapped, err := MapToResponse(order)
		if err != nil {
			return nil, fmt.Errorf("can't decode MapToResponses. Reason: %w", err)
		}
		orderResponses = append(orderResponses, *mapped)
	}
	return orderResponses, nil
}

func MapToPaginatedResponse(
	page global.PaginatedEntity[entities.Order],
) (*responses2.PaginatedResponse[responses.OrderResponse], error) {
	appResponses, err := MapToResponses(page.Content)
	if err != nil {
		return nil, fmt.Errorf("can't decode MapToPaginatedResponse. Reason: %w", err)
	}
	paginated := responses2.NewPaginatedResponse(appResponses, page.Pagination)
	return &paginated, nil
}
