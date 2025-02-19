// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"os"

	"github.com/SOAT-46/fastfood-operations/internal/orders"
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/gateways"
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/http/controllers"
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/repositories"
	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/implementations"
	"github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
	"github.com/SOAT-46/fastfood-operations/internal/shared/infrastructure/configuration"

	_ "github.com/SOAT-46/fastfood-operations/cmd/docs"
)

// Injectors from wire.go:

func injectApps() []entities.App {
	databaseSettings := newDatabaseSettings()
	database := configuration.MongoClient(databaseSettings)
	gormOrdersRepository := repositories.NewGormOrdersRepository(database)
	saveOrderGateway := gateways.NewSaveOrderGateway(gormOrdersRepository)
	createOrderUseCase := implementations.NewCreateOrderUseCase(saveOrderGateway)
	createOrderController := controllers.NewCreateOrderController(createOrderUseCase)
	getOrdersGateway := gateways.NewGetOrdersGateway(gormOrdersRepository)
	getOrdersUseCase := implementations.NewGetOrdersUseCase(getOrdersGateway)
	getOrdersController := controllers.NewGetOrdersController(getOrdersUseCase)
	updateOrderGateway := gateways.NewUpdateOrderGateway(gormOrdersRepository)
	getOrderByIDGateway := gateways.NewGetOrderByIDGateway(gormOrdersRepository)
	getOrderByIDUseCase := implementations.NewGetOrderByIDUseCase(getOrderByIDGateway)
	updateOrderUseCase := implementations.NewUpdateOrderUseCase(updateOrderGateway, getOrderByIDUseCase)
	updateOrderController := controllers.NewUpdateOrderController(updateOrderUseCase)
	getOrderByIDController := controllers.NewGetOrderByIDController(getOrderByIDUseCase)
	app := orders.NewApp(createOrderController, getOrdersController, updateOrderController, getOrderByIDController)
	v := newApps(app)
	return v
}

// wire.go:

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

func newApps(orders2 *orders.App,
) []entities.App {
	return []entities.App{orders2}
}
