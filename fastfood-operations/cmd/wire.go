//go:build !test && wireinject

package main

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders"
	"github.com/SOAT-46/fastfood-operations/internal/orders/infrastructure"
	"github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
	"github.com/SOAT-46/fastfood-operations/internal/shared/infrastructure/configuration"
	"github.com/google/wire"
	"os"
)

func injectApps() []entities.App {
	wire.Build(
		newDatabaseSettings,
		configuration.GormDB,
		infrastructure.Container,
		newApps,
	)
	return nil
}

func newDatabaseSettings() *entities.DatabaseSettings {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")
	ssl := os.Getenv("DB_SSL")
	return entities.NewDatabaseSettings(
		host,
		port,
		user,
		password,
		database,
		ssl,
	)
}

func newApps(
	orders *orders.App,
) []entities.App {
	return []entities.App{
		orders,
	}
}
