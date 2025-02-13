package models

import (
	"time"

	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoOrder struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty" mapstructure:"id"`
	Number     string             `bson:"number" json:"number" mapstructure:"number"`
	Status     string             `bson:"status" json:"status" mapstructure:"status"`
	ReceivedAt *time.Time         `bson:"received_at,omitempty" json:"received_at,omitempty" mapstructure:"received_at"`
	UpdatedAt  *time.Time         `bson:"updated_at,omitempty" json:"updated_at,omitempty" mapstructure:"updated_at"`
	Items      []MongoOrderItem   `bson:"items" json:"items" mapstructure:"items"`
}

func (itself MongoOrder) ToDomain() *entities.Order {
	status := entities.ToOrderStatus(itself.Status)
	return &entities.Order{
		Number:     itself.Number,
		Status:     status,
		ReceivedAt: itself.ReceivedAt,
		UpdatedAt:  itself.UpdatedAt,
	}
}

func BuildOrderFromDomain(entity entities.Order) MongoOrder {
	items := BuildGormOrderItemsFromDomain(entity.Items)
	return MongoOrder{
		ID:         primitive.NewObjectID(),
		Number:     entity.Number,
		Status:     entity.Status.String(),
		ReceivedAt: entity.ReceivedAt,
		UpdatedAt:  entity.UpdatedAt,
		Items:      items,
	}
}

func BuildOrderFromDomainInput(entity entities.CreateOrderInput) MongoOrder {
	items := BuildGormOrderItemsFromDomain(entity.Products)
	receivedAt := time.Now()
	return MongoOrder{
		ID:         primitive.NewObjectID(),
		Number:     entity.Number,
		Status:     entities.OrderStatusReceived.String(),
		ReceivedAt: &receivedAt,
		Items:      items,
	}
}
