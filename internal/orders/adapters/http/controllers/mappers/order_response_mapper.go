package mappers

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/http/responses"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	global "github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
	responses2 "github.com/SOAT-46/fastfood-operations/internal/shared/infrastructure/controllers/page"
)

func mapItems(products []entities.OrderProduct) []responses.OrderItemResponse {
	var items []responses.OrderItemResponse
	for _, product := range products {
		orderProduct := responses.OrderItemResponse{
			Product:  product.Product,
			Quantity: product.Quantity,
		}
		items = append(items, orderProduct)
	}
	return items
}

func MapToResponse(order entities.Order) *responses.OrderResponse {
	return &responses.OrderResponse{
		Number:     order.Number,
		Status:     order.Status.String(),
		ReceivedAt: order.ReceivedAt,
		Items:      mapItems(order.Items),
	}
}

func MapToResponses(orders []entities.Order) []responses.OrderResponse {
	var orderResponses []responses.OrderResponse
	for _, order := range orders {
		mapped := MapToResponse(order)
		orderResponses = append(orderResponses, *mapped)
	}
	return orderResponses
}

func MapToPaginatedResponse(
	page global.PaginatedEntity[entities.Order],
) *responses2.PaginatedResponse[responses.OrderResponse] {
	appResponses := MapToResponses(page.Content)
	paginated := responses2.NewPaginatedResponse(appResponses, page.Pagination)
	return &paginated
}
