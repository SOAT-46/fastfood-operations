package controllers

import "github.com/google/wire"

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	NewCreateOrderController,
	NewGetOrdersController,
	NewGetOrderByIDController,
	NewUpdateOrderController,
)
