package entities

type OrderProduct struct {
  Quantity  int
  ProductID int
}

func (itself OrderProduct) IsValid() bool {
  return itself.hasValidProduct() && itself.hasValidQuantity()
}

func (itself OrderProduct) hasValidProduct() bool {
  return itself.ProductID > 0
}

func (itself OrderProduct) hasValidQuantity() bool {
  return itself.Quantity > 0
}
