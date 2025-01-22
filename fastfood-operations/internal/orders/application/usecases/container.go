package usecases

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/contracts"
	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/implementations"
	"github.com/google/wire"
)

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	implementations.NewCreateOrderUseCase,
	wire.Bind(new(contracts.CreateOrder), new(*implementations.CreateOrderUseCase)),
	implementations.NewGetOrderByIDUseCase,
	wire.Bind(new(contracts.GetOrderByID), new(*implementations.GetOrderByIDUseCase)),
	implementations.NewGetOrdersUseCase,
	wire.Bind(new(contracts.GetOrders), new(*implementations.GetOrdersUseCase)),
	implementations.NewUpdateOrderUseCase,
	wire.Bind(new(contracts.UpdateOrder), new(*implementations.UpdateOrderUseCase)),
)
