package controllers

import (
	"net/http"

	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/http/helpers"
	_ "github.com/SOAT-46/fastfood-operations/internal/orders/adapters/http/responses" // for swagger
	"github.com/SOAT-46/fastfood-operations/internal/orders/application/usecases/contracts"
	entities2 "github.com/SOAT-46/fastfood-operations/internal/orders/domain/entities"
	"github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
	"github.com/SOAT-46/fastfood-operations/internal/shared/infrastructure/controllers/httperrors"
	_ "github.com/SOAT-46/fastfood-operations/internal/shared/infrastructure/controllers/page" // for swagger
	"github.com/SOAT-46/fastfood-operations/internal/shared/infrastructure/controllers/query"
	"github.com/gin-gonic/gin"
)

type GetOrdersController struct {
	gcontext *gin.Context
	useCase  contracts.GetOrders
}

func NewGetOrdersController(useCase contracts.GetOrders) *GetOrdersController {
	return &GetOrdersController{
		useCase: useCase,
	}
}

func (itself GetOrdersController) GetBind() entities.ControllerBind {
	return entities.ControllerBind{
		Method:       http.MethodGet,
		Version:      "v1",
		RelativePath: "/orders",
	}
}

// Execute Get orders
//
// @Summary Get orders
// @Description Get orders
// @BasePath /v1/orders
// @Tags orders
// @Accept application/json
// @Produce application/json
// @Param page query string false "page value"
// @Param size query string false "size value"
// @Param filter query string false	"filter value"
// @Success 200 {object} page.PaginatedResponse[responses.OrderResponse] "OK"
// @Failure 500 {object} httperrors.ErrorResponse "Internal Server Error"
// @Router /v1/orders [get].
func (itself GetOrdersController) Execute(gcontext *gin.Context) {
	itself.gcontext = gcontext
	pagination := query.GetPagination(itself.gcontext)

	listeners := contracts.GetOrdersListeners{
		OnSuccess: itself.onSuccess,
		OnError:   itself.onError,
	}
	itself.useCase.Execute(pagination, listeners)
}

func (itself GetOrdersController) onSuccess(page entities.PaginatedEntity[entities2.Order]) {
	helpers.HandlePaginatedResponse(itself.gcontext, http.StatusOK, page, itself.onError)
}

func (itself GetOrdersController) onError(err error) {
	response := httperrors.NewErrorResponse(err)
	itself.gcontext.JSON(http.StatusInternalServerError, response)
}
