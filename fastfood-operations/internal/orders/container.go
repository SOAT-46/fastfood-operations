package orders

import (
  "github.com/SOAT-46/fastfood-operations/internal/orders/adapters/http/controllers"
  "github.com/google/wire"
)

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
  controllers.Container,
  NewApp,
)
