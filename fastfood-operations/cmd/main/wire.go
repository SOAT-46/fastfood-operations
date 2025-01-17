//go:build !test && wireinject

package main

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders"
	"github.com/SOAT-46/fastfood-operations/internal/shared"
	"github.com/google/wire"
)

func injectApps() []shared.App {
	wire.Build(
		orders.Container,
		newApps,
	)
	return nil
}

func newApps(
	orders *orders.App,
) []shared.App {

	return []shared.App{
		orders,
	}
}
