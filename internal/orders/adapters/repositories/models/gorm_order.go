package models

import (
	"fmt"
	"time"

	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

type GormOrder struct {
	ID         int             `gorm:"primary_key;autoIncrement" mapstructure:"id"`
	Number     string          `gorm:"type:varchar(255);not null" mapstructure:"number"`
	Status     string          `gorm:"type:varchar(40);not null" mapstructure:"status"`
	ReceivedAt *time.Time      `gorm:"not null" mapstructure:"received_at"`
	UpdatedAt  *time.Time      `gorm:"default:null" mapstructure:"updated_at"`
	PaymentID  string          `gorm:"type:varchar(255);not null" mapstructure:"payment_id"`
	UserID     *string         `gorm:"type:varchar(255);default:null" mapstructure:"user_id"`
	Items      []GormOrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE" mapstructure:"items"`
}

func (GormOrder) TableName() string {
	return "orders"
}

func (itself GormOrder) ToDomain() (*entities.Order, error) {
	var response entities.Order
	err := mapstructure.Decode(itself, &response)
	if err != nil {
		return nil, fmt.Errorf("can't decode MapToResponse. Reason: %w", err)
	}
	response.ReceivedAt = itself.ReceivedAt
	response.UpdatedAt = itself.UpdatedAt

	return &response, nil
}

func BuildGormOrderFromDomain(entity entities.Order) GormOrder {
	items := BuildGormOrderItemsFromDomain(entity.Items)
	return GormOrder{
		ID:         entity.ID,
		Number:     entity.Number,
		Status:     entity.Status.String(),
		ReceivedAt: entity.ReceivedAt,
		UpdatedAt:  entity.UpdatedAt,
		PaymentID:  entity.Payment,
		UserID:     entity.UserID,
		Items:      items,
	}
}

func BuildGormOrderFromDomainInput(entity entities.CreateOrderInput) GormOrder {
	items := BuildGormOrderItemsFromDomain(entity.Products)
	number := uuid.New()
	receivedAt := time.Now()
	return GormOrder{
		Number:     number.String(),
		Status:     entities.OrderStatusReceived.String(),
		ReceivedAt: &receivedAt,
		PaymentID:  entity.PaymentID,
		UserID:     entity.UserID,
		Items:      items,
	}
}
