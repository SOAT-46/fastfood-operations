//go:build !test && wireinject

package main

import (
	"os"

	"github.com/SOAT-46/fastfood-operations/internal/orders"
	"github.com/SOAT-46/fastfood-operations/internal/orders/infrastructure"
	"github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
	"github.com/SOAT-46/fastfood-operations/internal/shared/infrastructure/configuration"
	"github.com/google/wire"
)

func injectApps() []entities.App {
	wire.Build(
		newDatabaseSettings,
		configuration.MongoClient,
		infrastructure.Container,
		newApps,
	)
	return nil
}

func newDatabaseSettings() *entities.DatabaseSettings {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	return entities.NewDatabaseSettings(
		host,
		port,
		user,
		password,
	)
}

func newApps(
	orders *orders.App,
) []entities.App {
	return []entities.App{
		orders,
	}
}
