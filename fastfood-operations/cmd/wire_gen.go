// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders"
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/gateways"
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/http/controllers"
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/repositories"
	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/implementations"
	"github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
	"github.com/SOAT-46/fastfood-operations/internal/shared/infrastructure/configuration"
	"os"
)

// Injectors from wire.go:

func injectApps() []entities.App {
	databaseSettings := newDatabaseSettings()
	db := configuration.GormDB(databaseSettings)
	gormOrdersRepository := repositories.NewGormOrdersRepository(db)
	saveOrderGateway := gateways.NewSaveOrderGateway(gormOrdersRepository)
	createOrderUseCase := implementations.NewCreateOrderUseCase(saveOrderGateway)
	createOrderController := controllers.NewCreateOrderController(createOrderUseCase)
	getOrdersGateway := gateways.NewGetOrdersGateway(gormOrdersRepository)
	getOrdersUseCase := implementations.NewGetOrdersUseCase(getOrdersGateway)
	getOrdersController := controllers.NewGetOrdersController(getOrdersUseCase)
	updateOrderGateway := gateways.NewUpdateOrderGateway(gormOrdersRepository)
	getOrderByIDGateway := gateways.NewGetOrderByIDGateway(gormOrdersRepository)
	updateOrderUseCase := implementations.NewUpdateOrderUseCase(updateOrderGateway, getOrderByIDGateway)
	updateOrderController := controllers.NewUpdateOrderController(updateOrderUseCase)
	getOrderByIDUseCase := implementations.NewGetOrderByIDUseCase(getOrderByIDGateway)
	getOrderByIDController := controllers.NewGetOrderByIDController(getOrderByIDUseCase)
	app := orders.NewApp(createOrderController, getOrdersController, updateOrderController, getOrderByIDController)
	v := newApps(app)
	return v
}

// wire.go:

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

func newApps(orders2 *orders.App,
) []entities.App {
	return []entities.App{orders2}
}
