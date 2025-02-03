package responses

import "time"

type OrderResponse struct {
	ID         int                 `json:"id,omitempty" mapstructure:"id"`
	Number     string              `json:"number,omitempty" mapstructure:"number"`
	Status     string              `json:"status,omitempty" mapstructure:"status"`
	ReceivedAt *time.Time          `json:"receivedAt,omitempty" mapstructure:"received_at"`
	Items      []OrderItemResponse `json:"items,omitempty" mapstructure:"items"`
} // @name OrderResponse

type OrderItemResponse struct {
	ProductID int `json:"product" mapstructure:"product_id"`
	Quantity  int `json:"quantity" mapstructure:"quantity"`
} // @name OrderItemResponse
