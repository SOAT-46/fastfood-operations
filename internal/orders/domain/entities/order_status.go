package entities

type OrderStatus string

const (
	OrderStatusPending     OrderStatus = "PENDING"
	OrderStatusReceived    OrderStatus = "RECEIVED"
	OrderStatusPreparation OrderStatus = "PREPARATION"
	OrderStatusReady       OrderStatus = "READY"
	OrderStatusDelivered   OrderStatus = "DELIVERED"
	OrderStatusCancelled   OrderStatus = "CANCELLED"
)

func (itself OrderStatus) String() string {
	return string(itself)
}

func ToOrderStatus(statusString string) OrderStatus {
	switch statusString {
	case string(OrderStatusPending):
		return OrderStatusPending
	case string(OrderStatusReceived):
		return OrderStatusReceived
	case string(OrderStatusPreparation):
		return OrderStatusPreparation
	case string(OrderStatusReady):
		return OrderStatusReady
	case string(OrderStatusDelivered):
		return OrderStatusDelivered
	case string(OrderStatusCancelled):
		return OrderStatusCancelled
	default:
		return ""
	}
}
