package requests

type CreateOrderRequest struct {
	Products []ProductRequest
	UserID   *int
}
