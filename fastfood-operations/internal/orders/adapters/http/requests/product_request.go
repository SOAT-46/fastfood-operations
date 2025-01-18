package requests

type ProductRequest struct {
	Quantity  int `json:"quantity"`
	OrderId   int `json:"order_id"`
	ProductId int `json:"product_id"`
}
