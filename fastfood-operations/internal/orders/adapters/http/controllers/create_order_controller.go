package controllers

import (
  "net/http"

  "github.com/SOAT-46/fastfood-operations/internal/shared"
  "github.com/gin-gonic/gin"
)

type CreateOrderController struct {
  gcontext *gin.Context
}

func NewCreateOrderController() *CreateOrderController {
  return &CreateOrderController{}
}

func (itself CreateOrderController) GetBind() shared.ControllerBind {
  return shared.ControllerBind{
    Method:       http.MethodPost,
    Version:      "v1",
    RelativePath: "/orders",
  }
}

// Execute Insert a new Order
// @Summary Insert a new Order
// @Description Insert a new Order
// @BasePath /v1/orders
// @Tags orders
// @Accept application/json
// @Produce application/json
// @Success 201
// @Router /v1/orders [post].
func (itself CreateOrderController) Execute(gcontext *gin.Context) {
  itself.gcontext = gcontext
  gcontext.JSON(http.StatusCreated, gin.H{
    "message": "ok",
  })
}
