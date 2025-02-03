package entities

import (
	"time"
)

type Order struct {
	ID         int            `mapstructure:"id"`
	Number     string         `mapstructure:"number"`
	Status     OrderStatus    `mapstructure:"status"`
	ReceivedAt *time.Time     `mapstructure:"received_at"`
	UpdatedAt  *time.Time     `mapstructure:"updated_at"`
	Payment    string         `mapstructure:"payment_id"`
	UserID     *string        `mapstructure:"user_id"`
	Items      []OrderProduct `mapstructure:"items"`
}

func (itself Order) IsValid() bool {
	return itself.hasValidNumber() && itself.hasValidStatus()
}

func (itself Order) hasValidNumber() bool {
	return len(itself.Number) > 0
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
