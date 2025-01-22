package builders

import (
	"bytes"
	"encoding/json"

	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/http/requests"
	"github.com/SOAT-46/fastfood-operations/test/shared/builders"
)

type UpdateOrderRequestBuilder struct {
	builders.BaseBuilder[requests.UpdateOrderRequest]
}

func NewUpdateOrderRequestBuilder() *UpdateOrderRequestBuilder {
	return &UpdateOrderRequestBuilder{}
}

func (itself *UpdateOrderRequestBuilder) BuildRequest() *bytes.Buffer {
	data := itself.Build()
	requestBodyBytes, _ := json.Marshal(data)
	return bytes.NewBuffer(requestBodyBytes)
}
