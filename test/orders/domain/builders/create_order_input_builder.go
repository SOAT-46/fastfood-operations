package builders

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	"github.com/SOAT-46/fastfood-operations/test/shared/builders"
)

type CreateOrderInputBuilder struct {
	builders.BaseBuilder[entities.CreateOrderInput]
}

func NewCreateOrderInputBuilder() *CreateOrderInputBuilder {
	return &CreateOrderInputBuilder{}
}

func (itself *CreateOrderInputBuilder) BuildInvalid() entities.CreateOrderInput {
	orderProduct := entities.OrderProduct{
		Quantity: 0,
		Product:  "",
	}
	return entities.CreateOrderInput{
		Products: []entities.OrderProduct{orderProduct},
	}
}
