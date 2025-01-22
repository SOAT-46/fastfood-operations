package models

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
)

type GormOrderItem struct {
	ID        int `gorm:"primary_key;autoIncrement" mapstructure:"id"`
	OrderID   int `gorm:"index;not null" mapstructure:"order_id"`
	ProductID int `gorm:"not null" mapstructure:"product_id"`
	Quantity  int `gorm:"not null" mapstructure:"quantity"`
}

func (GormOrderItem) TableName() string {
	return "order_product"
}

func BuildGormOrderItemFromDomain(entity entities.OrderProduct) GormOrderItem {
	return GormOrderItem{
		ID:        entity.ID,
		ProductID: entity.ProductID,
		Quantity:  entity.Quantity,
	}
}

func BuildGormOrderItemsFromDomain(entities []entities.OrderProduct) []GormOrderItem {
	items := make([]GormOrderItem, len(entities))
	for i, entity := range entities {
		items[i] = BuildGormOrderItemFromDomain(entity)
	}
	return items
}
