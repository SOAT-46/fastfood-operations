package entities

type OrderItem struct {
	Id        int `json:"id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
