package models

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoOrderItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty" mapstructure:"id"`
	Product  string             `bson:"product" json:"product" mapstructure:"product"`
	Quantity int                `bson:"quantity" json:"quantity" mapstructure:"quantity"`
}

func BuildGormOrderItemFromDomain(entity entities.OrderProduct) MongoOrderItem {
	return MongoOrderItem{
		ID:       primitive.NewObjectID(),
		Product:  entity.Product,
		Quantity: entity.Quantity,
	}
}

func BuildGormOrderItemsFromDomain(entities []entities.OrderProduct) []MongoOrderItem {
	items := make([]MongoOrderItem, len(entities))
	for i, entity := range entities {
		items[i] = BuildGormOrderItemFromDomain(entity)
	}
	return items
}
