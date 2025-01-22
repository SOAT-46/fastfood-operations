package models

type GormOrderItem struct {
	ID        string
	OrderID   string
	ProductID string
	Quantity  int
}

func (GormOrderItem) TableName() string {
	return "order_items"
}
