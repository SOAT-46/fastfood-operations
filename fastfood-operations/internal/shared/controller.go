package shared

import "github.com/gin-gonic/gin"

type Controller interface {
	GetBind() ControllerBind
	Execute(gcontext *gin.Context)
}
