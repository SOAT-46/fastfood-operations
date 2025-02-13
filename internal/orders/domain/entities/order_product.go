package entities

type OrderProduct struct {
	Product  string `mapstructure:"product"`
	Quantity int    `mapstructure:"quantity"`
}

func (itself OrderProduct) IsValid() bool {
	return itself.hasValidProduct() && itself.hasValidQuantity()
}

func (itself OrderProduct) hasValidProduct() bool {
	return itself.Product != ""
}

func (itself OrderProduct) hasValidQuantity() bool {
	return itself.Quantity > 0
}
