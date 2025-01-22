package entities

import (
	"time"
)

type Order struct {
	ID         int
	Number     int
	Status     OrderStatus
	ReceivedAt time.Time
	UpdatedAt  time.Time
	Payment    string
	UserID     *int
	Items      []OrderItem
}

func (itself Order) IsValid() bool {
	return itself.hasValidNumber() && itself.hasValidStatus()
}

func (itself Order) hasValidNumber() bool {
	return itself.Number > 0
}

func (itself Order) hasValidStatus() bool {
	validStatuses := []OrderStatus{
		OrderStatusPending,
		OrderStatusReceived,
		OrderStatusPreparation,
		OrderStatusDelivered,
		OrderStatusReady,
		OrderStatusCancelled,
	}

	for _, status := range validStatuses {
		if itself.Status == status {
			return true
		}
	}
	return false
}
