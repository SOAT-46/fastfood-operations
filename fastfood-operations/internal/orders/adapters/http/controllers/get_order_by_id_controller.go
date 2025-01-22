package controllers

import (
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/http/helpers"
	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/contracts"
	entities2 "github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	"github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
	"github.com/SOAT-46/fastfood-operations/internal/shared/infrastructure/controllers/httperrors"
	"github.com/SOAT-46/fastfood-operations/internal/shared/infrastructure/controllers/query"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetOrderByIDController struct {
	gcontext *gin.Context
	useCase  contracts.GetOrderByID
}

func NewGetOrderByIDController(useCase contracts.GetOrderByID) *GetOrderByIDController {
	return &GetOrderByIDController{
		useCase: useCase,
	}
}

func (itself GetOrderByIDController) GetBind() entities.ControllerBind {
	return entities.ControllerBind{
		Method:       http.MethodGet,
		Version:      "v1",
		RelativePath: "/orders/:id",
	}
}

// Execute Get an order by the target ID
//
// @Summary Get an order by the target ID
// @Description Get an order by the target ID
// @BasePath /v1/orders/:id
// @Tags orders
// @Accept application/json
// @Produce application/json
// @Param id path int64 true "Order ID"
// @Success 200 {object} responses.OrderResponse "OK"
// @Failure 404 {object} httperrors.ErrorResponse "Not Found"
// @Failure 500 {object} httperrors.ErrorResponse "Internal Server Error"
// @Router /v1/orders/{id} [get].
func (itself GetOrderByIDController) Execute(gcontext *gin.Context) {
	itself.gcontext = gcontext
	id := query.GetID(itself.gcontext)

	listeners := contracts.GetOrderByIDListeners{
		OnSuccess:  itself.onSuccess,
		OnNotFound: itself.onNotFound,
		OnError:    itself.onError,
	}
	itself.useCase.Execute(id, listeners)
}

func (itself GetOrderByIDController) onSuccess(order entities2.Order) {
	helpers.HandleResponse(itself.gcontext, http.StatusOK, order, itself.onError)
}

func (itself GetOrderByIDController) onNotFound() {
	itself.gcontext.JSON(http.StatusNotFound, gin.H{
		"message": "Order not found",
	})
}

func (itself GetOrderByIDController) onError(err error) {
	response := httperrors.NewErrorResponse(err)
	itself.gcontext.JSON(http.StatusInternalServerError, response)
}
