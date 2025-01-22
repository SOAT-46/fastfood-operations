package responses

import "time"

type OrderResponse struct {
	Number     int                 `json:"number"`
	Status     string              `json:"status"`
	ReceivedAt time.Time           `json:"receivedAt"`
	Items      []OrderItemResponse `json:"items"`
} // @name OrderResponse

type OrderItemResponse struct {
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
} // @name OrderItemResponse
