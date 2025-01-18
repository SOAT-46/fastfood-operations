package orders

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/http/controllers"
	"github.com/SOAT-46/fastfood-operations/internal/shared"
)

type App struct {
	controllers []shared.Controller
}

func NewApp(createController *controllers.CreateOrderController) *App {
	appControllers := []shared.Controller{
		createController,
	}
	return &App{appControllers}
}

func (itself *App) GetControllers() []shared.Controller {
	return itself.controllers
}
