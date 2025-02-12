package query

import (
	"github.com/gin-gonic/gin"
)

const ID string = "id"

func GetID(ctx *gin.Context) string {
	return getParam(ctx, ID)
}

func getParam(ctx *gin.Context, name string) string {
	return ctx.Param(name)
}
