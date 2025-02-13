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
	"github.com/gin-gonic/gin"
)

type CreateOrderController struct {
	gcontext *gin.Context
	useCase  contracts.CreateOrder
}

func NewCreateOrderController(useCase contracts.CreateOrder) *CreateOrderController {
	return &CreateOrderController{
		useCase: useCase,
	}
}

func (itself CreateOrderController) GetBind() entities.ControllerBind {
	return entities.ControllerBind{
		Method:       http.MethodPost,
		Version:      "v1",
		RelativePath: "/orders",
	}
}

// Execute Insert a new Order
//
// @Summary Insert a new Order
// @Description Insert a new Order
// @BasePath /v1/orders
// @Tags orders
// @Accept application/json
// @Produce application/json
// @Param request body requests.CreateOrderRequest true "Request body"
// @Success 201 {object} responses.OrderResponse "Created"
// @Failure 400 {object} httperrors.ErrorResponse "Bad Request"
// @Failure 500 {object} httperrors.ErrorResponse "Internal Server Error"
// @Router /v1/orders [post].
func (itself CreateOrderController) Execute(gcontext *gin.Context) {
	itself.gcontext = gcontext
	ctx, cancel := controllers.DefaultTimeout(itself.gcontext)
	defer cancel()

	var body requests.CreateOrderRequest
	if err := gcontext.ShouldBindJSON(&body); err != nil {
		itself.onInvalid(err)
		return
	}

	listeners := contracts.CreateOrderListeners{
		OnSuccess: itself.onSuccess,
		OnInvalid: itself.onInvalid,
		OnError:   itself.onError,
	}

	input := body.ToDomain()
	itself.useCase.Execute(ctx, input, listeners)
}

func (itself CreateOrderController) onSuccess(order entities2.Order) {
	helpers.HandleResponse(itself.gcontext, http.StatusCreated, order)
}

func (itself CreateOrderController) onInvalid(err error) {
	response := httperrors.NewErrorResponse(err)
	itself.gcontext.JSON(http.StatusBadRequest, response)
}

func (itself CreateOrderController) onError(err error) {
	response := httperrors.NewErrorResponse(err)
	itself.gcontext.JSON(http.StatusInternalServerError, response)
}
