package controllers

import (
	"net/http"

	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/http/controllers/helpers"
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/http/requests"
	_ "github.com/SOAT-46/fastfood-operations/internal/orders/adapters/http/responses" // for swagger
	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/contracts"
	entities2 "github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	"github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
	"github.com/SOAT-46/fastfood-operations/internal/shared/infrastructure/controllers"
	"github.com/SOAT-46/fastfood-operations/internal/shared/infrastructure/controllers/httperrors"
	"github.com/SOAT-46/fastfood-operations/internal/shared/infrastructure/controllers/query"
	"github.com/gin-gonic/gin"
)

type UpdateOrderController struct {
	gcontext *gin.Context
	useCase  contracts.UpdateOrder
}

func NewUpdateOrderController(useCase contracts.UpdateOrder) *UpdateOrderController {
	return &UpdateOrderController{
		useCase: useCase,
	}
}

func (itself UpdateOrderController) GetBind() entities.ControllerBind {
	return entities.ControllerBind{
		Method:       http.MethodPut,
		Version:      "v1",
		RelativePath: "/orders/:id",
	}
}

// Execute Update an order
//
// @Summary Update an order
// @Description Update an order
// @BasePath /v1/orders/:id
// @Tags orders
// @Accept application/json
// @Produce application/json
// @Param id path string true "Order ID"
// @Param request body requests.UpdateOrderRequest true "Request body"
// @Success 200 {object} responses.OrderResponse "OK"
// @Failure 404 {object} httperrors.ErrorResponse "Not Found"
// @Failure 500 {object} httperrors.ErrorResponse "Internal Server Error"
// @Router /v1/orders/{id} [put].
func (itself UpdateOrderController) Execute(gcontext *gin.Context) {
	itself.gcontext = gcontext
	id := query.GetID(itself.gcontext)

	ctx, cancel := controllers.DefaultTimeout(itself.gcontext)
	defer cancel()

	var body requests.UpdateOrderRequest
	if err := gcontext.ShouldBindJSON(&body); err != nil {
		itself.onInvalid(err)
		return
	}

	listeners := contracts.UpdateOrderListeners{
		OnSuccess:  itself.onSuccess,
		OnNotFound: itself.onNotFound,
		OnError:    itself.onError,
	}
	order := body.ToDomain(id)
	itself.useCase.Execute(ctx, order, listeners)
}

func (itself UpdateOrderController) onSuccess(order entities2.Order) {
	helpers.HandleResponse(itself.gcontext, http.StatusOK, order)
}

func (itself UpdateOrderController) onNotFound() {
	itself.gcontext.JSON(http.StatusNotFound, gin.H{
		"error": "Not found answer",
	})
}

func (itself UpdateOrderController) onInvalid(err error) {
	response := httperrors.NewErrorResponse(err)
	itself.gcontext.JSON(http.StatusBadRequest, response)
}

func (itself UpdateOrderController) onError(err error) {
	response := httperrors.NewErrorResponse(err)
	itself.gcontext.JSON(http.StatusInternalServerError, response)
}
