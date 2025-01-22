package orders

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/http/controllers"
	"github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
)

type App struct {
	controllers []entities.Controller
}

func NewApp(
	createController *controllers.CreateOrderController,
	getOrdersController *controllers.GetOrdersController,
	updateOrderController *controllers.UpdateOrderController,
	getOrderController *controllers.GetOrderByIDController,
) *App {
	appControllers := []entities.Controller{
		createController,
		getOrdersController,
		updateOrderController,
		getOrderController,
	}
	return &App{appControllers}
}

func (itself *App) GetControllers() []entities.Controller {
	return itself.controllers
}
