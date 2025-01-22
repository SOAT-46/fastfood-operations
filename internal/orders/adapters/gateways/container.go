package gateways

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/gateways"
	"github.com/google/wire"
)

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	NewGetOrdersGateway,
	wire.Bind(new(gateways.GetOrdersPort), new(*GetOrdersGateway)),
	NewGetOrderByIDGateway,
	wire.Bind(new(gateways.GetOrderByIDPort), new(*GetOrderByIDGateway)),
	NewUpdateOrderGateway,
	wire.Bind(new(gateways.UpdateOrderPort), new(*UpdateOrderGateway)),
	NewSaveOrderGateway,
	wire.Bind(new(gateways.SaveOrderPort), new(*SaveOrderGateway)),
)
