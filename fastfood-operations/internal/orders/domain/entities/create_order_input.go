package entities

type CreateOrderInput struct {
	Products []OrderProduct
	UserID   *int
}

func (input *CreateOrderInput) IsValid() bool {
	return allProductsValid(input.Products)
}

func allProductsValid(products []OrderProduct) bool {
	for _, product := range products {
		if !product.IsValid() {
			return false
		}
	}
	return true
}
