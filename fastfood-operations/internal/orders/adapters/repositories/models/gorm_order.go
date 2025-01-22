package models

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	"time"
)

type GormOrder struct {
	ID         string
	Number     int64
	Status     string
	ReceivedAt time.Time
	UpdatedAt  *time.Time
	PaymentID  string
	UserID     *string
	Items      []GormOrderItem
}

func (GormOrder) TableName() string {
	return "orders"
}

func (itself GormOrder) ToDomain() *entities.Order {
	return &entities.Order{}
}

func BuildGormOrderFromDomain(entity entities.Order) GormOrder {
	return GormOrder{}
}

func BuildGormOrderFromDomainInput(entity entities.CreateOrderInput) GormOrder {
	return GormOrder{}
}
