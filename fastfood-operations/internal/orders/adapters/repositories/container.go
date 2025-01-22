package repositories

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/repositories/contracts"
	"github.com/google/wire"
)

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	NewGormOrdersRepository,
	wire.Bind(new(contracts.OrdersRepository), new(*GormOrdersRepository)),
)
