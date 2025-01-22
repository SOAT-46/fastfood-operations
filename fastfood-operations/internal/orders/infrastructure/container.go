package infrastructure

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders"
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/gateways"
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/http/controllers"
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/repositories"
	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases"
	"github.com/google/wire"
)

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	repositories.Container,
	gateways.Container,
	usecases.Container,
	controllers.Container,
	orders.NewApp,
)
