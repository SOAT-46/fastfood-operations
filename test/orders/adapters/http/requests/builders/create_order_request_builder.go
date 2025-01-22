package builders

import (
	"bytes"
	"encoding/json"

	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/http/requests"
	"github.com/SOAT-46/fastfood-operations/test/shared/builders"
)

type CreateOrderRequestBuilder struct {
	builders.BaseBuilder[requests.CreateOrderRequest]
}

func NewCreateOrderRequestBuilder() *CreateOrderRequestBuilder {
	return &CreateOrderRequestBuilder{}
}

func (itself *CreateOrderRequestBuilder) BuildRequest() *bytes.Buffer {
	data := itself.Build()
	requestBodyBytes, _ := json.Marshal(data)
	return bytes.NewBuffer(requestBodyBytes)
}
