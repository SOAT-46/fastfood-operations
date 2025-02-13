package entities

import (
	"time"
)

type Order struct {
	Number     string         `mapstructure:"number"`
	Status     OrderStatus    `mapstructure:"status"`
	ReceivedAt *time.Time     `mapstructure:"received_at"`
	UpdatedAt  *time.Time     `mapstructure:"updated_at"`
	Items      []OrderProduct `mapstructure:"items"`
}
