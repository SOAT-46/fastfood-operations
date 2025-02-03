package query

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

const ID string = "id"

func GetID(ctx *gin.Context) int {
	return getAsInt(ctx, ID)
}

func getAsInt(ctx *gin.Context, paramName string) int {
	value := getParam(ctx, paramName)
	param, _ := strconv.Atoi(value)

	return param
}

func getParam(ctx *gin.Context, name string) string {
	return ctx.Param(name)
}
